package dashboardbuilder

import (
	"context"
	"fmt"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"

	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"
)

const suiEmoji = "💧"

type DashboardBuilder struct {
	Ctx       context.Context
	Terminal  *termbox.Terminal
	Dashboard *container.Container
	Cells     []*Cell
	Quitter   func(k *terminalapi.Keyboard)
}

func NewDashboardBuilder() (*DashboardBuilder, error) {
	var (
		terminal  *termbox.Terminal
		dashboard *container.Container
		cells     []*Cell
		err       error
	)

	ctx, cancel := context.WithCancel(context.Background())

	cells = initCells()

	if terminal, err = termbox.New(); err != nil {
		return nil, err
	}

	if dashboard, err = initDashboard(terminal); err != nil {
		return nil, err
	}

	for idx, cell := range cells {
		cellNameEnum := enums.CellName(idx)
		cellID := cellNameEnum.String()

		dashboard.Update(cellID, cell.Config...)
	}

	return &DashboardBuilder{
		Ctx:       ctx,
		Terminal:  terminal,
		Dashboard: dashboard,
		Cells:     cells,
		Quitter: func(k *terminalapi.Keyboard) {
			if k.Key == 'q' || k.Key == 'Q' || k.Key == keyboard.KeyEsc {
				terminal.Close()
				cancel()
			}
		},
	}, nil
}

func initCells() []*Cell {
	var (
		cellNames = enums.CellNameValues()
		cells     = make([]*Cell, len(cellNames))
	)

	for _, name := range cellNames {
		var (
			cellName = name.CellNameString()
			cell     *Cell
		)

		if name == enums.CellNameCheckSyncProgress || name == enums.CellNameTXSyncProgress {
			cell = NewProgressCell(cellName)
		} else {
			cell = NewTextCell(cellName)
		}

		cells[name] = cell
	}

	return cells
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
										container.ID(enums.CellNameStatus.String()),
									),
									container.Right(
										container.ID(enums.CellNameAddress.String()),
									),
									container.SplitFixed(8),
								),
							),
							container.Bottom(
								container.SplitVertical(
									container.Left(
										container.ID(enums.CellNameVersion.String()),
									),
									container.Right(
										container.ID(enums.CellNameCommit.String()),
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
														container.ID(enums.CellNameLatestCheckpoint.String()),
													),
													container.Right(
														container.ID(enums.CellNameHighestCheckpoint.String()),
													),
												),
											),
											container.Right(
												container.SplitVertical(
													container.Left(
														container.ID(enums.CellNameConnectedPeers.String()),
													),
													container.Right(
														container.ID(enums.CellNameUptime.String()),
													),
												),
											),
										),
									),
									container.Right(
										container.ID(enums.CellNameTXSyncProgress.String()),
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
														container.ID(enums.CellNameTotalTransactions.String()),
													),
													container.Right(
														container.ID(enums.CellNameTransactionsPerSecond.String()),
													),
												),
											),
											container.Right(
												container.SplitVertical(
													container.Left(
														container.ID(enums.CellNameCountry.String()),
													),
													container.Right(
														container.ID(enums.CellNameCompany.String()),
													),
												),
											),
										),
									),
									container.Right(
										container.ID(enums.CellNameCheckSyncProgress.String()),
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
