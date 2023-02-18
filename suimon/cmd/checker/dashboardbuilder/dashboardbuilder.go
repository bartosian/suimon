package dashboardbuilder

import (
	"context"
	"fmt"

	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgetapi"
	"github.com/mum4k/termdash/widgets/gauge"
	"github.com/mum4k/termdash/widgets/text"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/dashboardbuilder/dashboards"
)

const suiEmoji = "💧"

func NewTextWidget() (*text.Text, error) {
	return text.New(text.RollContent(), text.WrapAtWords())
}

func NewProgressWidget() (*gauge.Gauge, error) {
	return gauge.New(
		gauge.Height(3),
		gauge.Border(linestyle.Light, cell.FgColor(cell.ColorGreen)),
		gauge.Color(cell.ColorGreen),
		gauge.FilledTextColor(cell.ColorBlack),
		gauge.EmptyTextColor(cell.ColorWhite),
		gauge.HorizontalTextAlign(align.HorizontalCenter),
	)
}

func NewCell(title string, widget widgetapi.Widget) []container.Option {
	return []container.Option{
		container.FocusedColor(cell.ColorGreen),
		container.Border(linestyle.Light),
		container.BorderTitle(title),
		container.PlaceWidget(widget),
		container.AlignVertical(align.VerticalMiddle),
		container.AlignHorizontal(align.HorizontalCenter),
		container.PaddingTopPercent(2),
		container.BorderColor(cell.ColorRed),
		container.TitleColor(cell.ColorGreen),
	}
}

type DashboardBuilder struct {
	Ctx       context.Context
	Terminal  *termbox.Terminal
	Dashboard *container.Container
	Widgets   []widgetapi.Widget
	Quitter   func(k *terminalapi.Keyboard)
}

func NewDashboardBuilder() (*DashboardBuilder, error) {
	var (
		terminal  *termbox.Terminal
		dashboard *container.Container
		widgets   []widgetapi.Widget
		err       error
	)

	ctx, cancel := context.WithCancel(context.Background())

	if terminal, err = termbox.New(); err != nil {
		return nil, err
	}

	if dashboard, err = initDashboard(terminal); err != nil {
		return nil, err
	}

	if widgets, err = initWidgets(); err != nil {
		return nil, err
	}

	for idx, widget := range widgets {
		cellNameEnum := dashboards.CellName(idx)
		cellName := cellNameEnum.CellNameString()
		cellID := cellNameEnum.String()
		cell := NewCell(cellName, widget)

		dashboard.Update(cellID, cell...)
	}

	return &DashboardBuilder{
		Ctx:       ctx,
		Terminal:  terminal,
		Dashboard: dashboard,
		Widgets:   widgets,
		Quitter: func(k *terminalapi.Keyboard) {
			if k.Key == 'q' || k.Key == 'Q' || k.Key == keyboard.KeyEsc {
				terminal.Close()
				cancel()
			}
		},
	}, nil
}

func initWidgets() ([]widgetapi.Widget, error) {
	var (
		totalTransactionsWidget *text.Text
		latestCheckpointWidget  *text.Text
		syncedCheckpointWidget  *text.Text
		connecetedPeersWidget   *text.Text
		statusWidget            *text.Text
		addressWidget           *text.Text
		versionWidget           *text.Text
		commitWidget            *text.Text
		tpsWidget               *text.Text
		uptimeWidget            *text.Text
		providerWidget          *text.Text
		countryWidget           *text.Text
		txSyncWidget            *gauge.Gauge
		checkSyncWidget         *gauge.Gauge
		err                     error
	)

	if statusWidget, err = NewTextWidget(); err != nil {
		return nil, err
	}

	if addressWidget, err = NewTextWidget(); err != nil {
		return nil, err
	}

	if tpsWidget, err = NewTextWidget(); err != nil {
		return nil, err
	}

	if totalTransactionsWidget, err = NewTextWidget(); err != nil {
		return nil, err
	}

	if latestCheckpointWidget, err = NewTextWidget(); err != nil {
		return nil, err
	}

	if syncedCheckpointWidget, err = NewTextWidget(); err != nil {
		return nil, err
	}

	if connecetedPeersWidget, err = NewTextWidget(); err != nil {
		return nil, err
	}

	if uptimeWidget, err = NewTextWidget(); err != nil {
		return nil, err
	}

	if versionWidget, err = NewTextWidget(); err != nil {
		return nil, err
	}

	if commitWidget, err = NewTextWidget(); err != nil {
		return nil, err
	}

	if providerWidget, err = NewTextWidget(); err != nil {
		return nil, err
	}

	if countryWidget, err = NewTextWidget(); err != nil {
		return nil, err
	}

	if txSyncWidget, err = NewProgressWidget(); err != nil {
		return nil, err
	}

	if checkSyncWidget, err = NewProgressWidget(); err != nil {
		return nil, err
	}

	return []widgetapi.Widget{
		dashboards.CellNameStatus:                statusWidget,
		dashboards.CellNameAddress:               addressWidget,
		dashboards.CellNameTransactionsPerSecond: tpsWidget,
		dashboards.CellNameTotalTransactions:     totalTransactionsWidget,
		dashboards.CellNameLatestCheckpoint:      latestCheckpointWidget,
		dashboards.CellNameHighestCheckpoint:     syncedCheckpointWidget,
		dashboards.CellNameConnectedPeers:        connecetedPeersWidget,
		dashboards.CellNameTXSyncProgress:        txSyncWidget,
		dashboards.CellNameCheckSyncProgress:     checkSyncWidget,
		dashboards.CellNameUptime:                uptimeWidget,
		dashboards.CellNameVersion:               versionWidget,
		dashboards.CellNameCommit:                commitWidget,
		dashboards.CellNameCompany:               providerWidget,
		dashboards.CellNameCountry:               countryWidget,
	}, nil
}

func initDashboard(terminal *termbox.Terminal) (*container.Container, error) {
	return container.New(
		terminal,
		container.Border(linestyle.Light),
		container.BorderColor(cell.ColorGreen),
		container.FocusedColor(cell.ColorGreen),
		container.BorderTitle(fmt.Sprintf("%s SUIMON: PRESS Q or ESC TO QUIT", suiEmoji)),
		container.PaddingTopPercent(1),
		container.Focused(),
		container.AlignHorizontal(align.HorizontalCenter),
		container.AlignVertical(align.VerticalMiddle),
		container.SplitHorizontal(
			container.Top(
				container.SplitVertical(
					container.Left(
						container.SplitHorizontal(
							container.Top(
								container.SplitVertical(
									container.Left(
										container.ID(dashboards.CellNameStatus.String()),
									),
									container.Right(
										container.ID(dashboards.CellNameAddress.String()),
									),
									container.SplitFixed(8),
								),
							),
							container.Bottom(
								container.SplitVertical(
									container.Left(
										container.ID(dashboards.CellNameVersion.String()),
									),
									container.Right(
										container.ID(dashboards.CellNameCommit.String()),
									),
								),
							),
						),
					),
					container.Right(
						container.SplitHorizontal(
							container.Top(
								container.SplitVertical(
									container.Left(
										container.SplitVertical(
											container.Left(
												container.SplitVertical(
													container.Left(
														container.ID(dashboards.CellNameLatestCheckpoint.String()),
													),
													container.Right(
														container.ID(dashboards.CellNameHighestCheckpoint.String()),
													),
												),
											),
											container.Right(
												container.SplitVertical(
													container.Left(
														container.ID(dashboards.CellNameConnectedPeers.String()),
													),
													container.Right(
														container.ID(dashboards.CellNameUptime.String()),
													),
												),
											),
										),
									),
									container.Right(
										container.ID(dashboards.CellNameTXSyncProgress.String()),
									),
								),
							),
							container.Bottom(
								container.SplitVertical(
									container.Left(
										container.SplitVertical(
											container.Left(
												container.SplitVertical(
													container.Left(
														container.ID(dashboards.CellNameTotalTransactions.String()),
													),
													container.Right(
														container.ID(dashboards.CellNameTransactionsPerSecond.String()),
													),
												),
											),
											container.Right(
												container.SplitVertical(
													container.Left(
														container.ID(dashboards.CellNameCountry.String()),
													),
													container.Right(
														container.ID(dashboards.CellNameCompany.String()),
													),
												),
											),
										),
									),
									container.Right(
										container.ID(dashboards.CellNameCheckSyncProgress.String()),
									),
								),
							),
						),
					),
					container.SplitPercent(15),
				),
			),
			container.Bottom(),
			container.SplitPercent(15),
		),
	)
}
