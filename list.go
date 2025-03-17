package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/pkg/errors"
)

const listHeight = 16

const defaultWidth = 80

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = DefaultStyle.PaddingLeft(2)
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

// InfoListItem is the interface for an info list. If an object satisfies this
// interface, a list prompt can be generated from a slice of these values.
// Additionally, the value returns from Info will render as markdown.
type InfoListItem interface {
	FilterValue() string
	GetName() string
	Info() string
}

// StringItem is a type that satisfies the InfoListItem interface. It's useful
// when you just want to display a list of strings.
type StringItem string

// FilterValue returns the StringItem's value, which is used for filtering.
func (s StringItem) FilterValue() string { return string(s) }

// GetName returns the StringItem's value, which is the value that is displayed
// in the list.
func (s StringItem) GetName() string { return string(s) }

// Info returns the StringItem's value, which is displayed when the user selects
// the detail view for the item.
func (s StringItem) Info() string { return string(s) }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(InfoListItem)
	if !ok {
		return
	}
	str := fmt.Sprintf("%d. %s", index+1, i.GetName())

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(strs ...string) string {
			return selectedItemStyle.Render("> " + strs[0])
		}
	}

	fmt.Fprint(w, fn(str))
}

// InfoListModel contains the state of the info list prompt.
type InfoListModel struct {
	list     list.Model
	vp       viewport.Model
	renderer *glamour.TermRenderer
	choice   InfoListItem
	quitting bool
	detail   bool

	disableDetail bool
}

// NewInfoListModelInput is the input for the NewInfoListModel function.
type NewInfoListModelInput struct {
	Title            string
	Items            []InfoListItem
	NameSingular     string
	NamePlural       string
	DisableFiltering bool
	DisableDetail    bool
}

// NewInfoListModel creates a new InfoListModel with defaults.
func NewInfoListModel(input NewInfoListModelInput) (InfoListModel, error) {
	listItems := make([]list.Item, len(input.Items))
	for i, item := range input.Items {
		listItems[i] = item
	}
	l := list.New(listItems, itemDelegate{}, defaultWidth, listHeight)
	if input.Title != "" {
		l.Title = input.Title
	}
	if input.NameSingular != "" && input.NamePlural != "" {
		l.SetStatusBarItemName(input.NameSingular, input.NamePlural)
		l.SetShowStatusBar(true)
	}
	l.SetFilteringEnabled(!input.DisableFiltering)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	if !input.DisableDetail {
		l.AdditionalShortHelpKeys = func() []key.Binding {
			return []key.Binding{
				key.NewBinding(
					key.WithKeys(" "),
					key.WithHelp("space", "show details"),
				),
			}
		}
	}

	vp := viewport.New(defaultWidth, listHeight-2)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(frgMagenta)).
		PaddingRight(2)

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(defaultWidth),
	)
	if err != nil {
		return InfoListModel{}, errors.Wrap(err, "failed to create markdown renderer")
	}
	return InfoListModel{
		list:          l,
		vp:            vp,
		renderer:      renderer,
		disableDetail: input.DisableDetail,
	}, nil
}

// Init is the first command that is called when the program starts.
func (m InfoListModel) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m InfoListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.detail {
	case true: // detail view
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "q", "ctrl+c":
				m.quitting = true
				return m, tea.Quit
			case " ", "esc":
				m.detail = !m.detail
				return m, nil
			default:
				var cmd tea.Cmd
				m.vp, cmd = m.vp.Update(msg)
				return m, cmd
			}
		default:
			return m, nil
		}
	default: // list view
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.list.SetWidth(msg.Width)
			return m, nil

		case tea.KeyMsg:
			switch keypress := msg.String(); keypress {
			case "ctrl+c":
				m.quitting = true
				return m, tea.Quit
			case "enter":
				i, ok := m.list.SelectedItem().(InfoListItem)
				if ok {
					m.choice = i
				}
				return m, tea.Quit
			case " ":
				if m.disableDetail {
					return m, nil
				}
				m.detail = !m.detail
				str, err := m.renderer.Render(m.list.SelectedItem().(InfoListItem).Info())
				if err != nil {
					m.vp.SetContent(fmt.Sprintf("failed to render detail view: %v", err))
				}
				m.vp.SetYOffset(0)
				m.vp.SetContent(str)
				return m, nil
			}
		}
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}
}

// View renders the model.
func (m InfoListModel) View() string {
	itemName, _ := m.list.StatusBarItemName()
	if m.choice != nil {
		return quitTextStyle.Render(fmt.Sprintf("%s %q selected", itemName, m.choice.GetName()))
	}
	if m.quitting {
		return quitTextStyle.Render(fmt.Sprintf("no %s selected", itemName))
	}
	if m.detail {
		return m.vp.View() + "\n" + m.viewportHelpView()
	}
	return "\n" + m.list.View()
}

// Run starts the prompt.
func (m InfoListModel) Run() (InfoListItem, error) {
	mFinal, err := tea.NewProgram(m).Run()
	if err != nil {
		return nil, errors.Wrap(err, "failed to start list program")
	}
	return mFinal.(InfoListModel).choice, nil
}

func (m InfoListModel) viewportHelpView() string {
	return m.list.Styles.HelpStyle.Copy().UnsetPaddingBottom().Render(
		m.list.Help.ShortHelpView([]key.Binding{
			key.NewBinding(
				key.WithKeys("up", "down"),
				key.WithHelp("↑/↓", "navigate"),
			),
			key.NewBinding(
				key.WithKeys(" ", "esc"),
				key.WithHelp("space/esc", "back to list"),
			),
			key.NewBinding(
				key.WithKeys("q"),
				key.WithHelp("q", "quit"),
			),
		})) +
		m.list.Styles.HelpStyle.Copy().UnsetPaddingTop().
			Render(fmt.Sprintf("scroll %3.f%%", m.vp.ScrollPercent()*100))
}
