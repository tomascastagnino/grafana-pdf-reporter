package services

import (
	"log"
	"net/http"
	"time"

	"github.com/tomascastagnino/grafana-pdf-reporter/internal/clients"
	"github.com/tomascastagnino/grafana-pdf-reporter/internal/models"
)

type DashboardService struct {
	grafanaClient clients.GrafanaClient
	panelService  *PanelService
}

func NewDashboardService(client clients.GrafanaClient, panelService *PanelService) *DashboardService {
	return &DashboardService{grafanaClient: client, panelService: panelService}
}

func (s *DashboardService) GetDashboard(dashboardID string, r *http.Request) (*models.Dashboard, error) {
	name := "GetDashboard"
	start := time.Now()
	log.Printf("Starting %s", name)
	// Fetch dashboard metadata
	dashboard, err := s.grafanaClient.GetDashboard(dashboardID, r.Header)
	if err != nil {
		return nil, err
	}
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)

	// Get panels with images using the PanelService
	panels, err := s.panelService.GetPanelsWithImages(dashboard, *r) // Note: You can pass *r here if needed.
	if err != nil {
		return nil, err
	}

	// Assign panels to the dashboard
	dashboard.Panels = panels

	return dashboard, nil
}

func (s *DashboardService) ListDashboards() ([]models.Dashboard, error) {
	panic("TODO: To be implemented in the future.")
}
