package manager

import "github.com/reonardoleis/cherry/pkg/page"

var manager *Manager

type Manager struct {
	activePage *page.Page
	pages      []*page.Page
}

func init() {
	manager = &Manager{}
}

func Instance() *Manager {
	return manager
}

func (m *Manager) RegisterPage(page *page.Page) {
	m.pages = append(m.pages, page)
}

func (m *Manager) ActivePage() *page.Page {
	if m.activePage == nil {
		return m.pages[0]
	}

	return m.activePage
}
