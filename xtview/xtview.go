package xtview

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"math/rand"
	"root"
)

const refreshInterval = 600 * time.Millisecond

var (
	message = []string{
		"Connecting to satellite",
		"Transferring Funds",
		"Defragmenting Drive",
		"Increasing XP",
		"Brewing Tea",
		"Charging Tachyons",
		"Deinterlacing Stream",
		"Increasing particle spallation",
		"Enriching Uranium",
		"Checking for bad sectors",
		"Saving boot sector",
		// Sim City
		"Adding Hidden Agendas",
		"Adjusting Bell Curves",
		"Aesthesizing Industrial Areas",
		"Aligning Covariance Matrices",
		"Applying Feng Shui Shaders",
		"Applying Theatre Soda Layer",
		"Asserting Packed Exemplars",
		"Attempting to Lock Back-Buffer",
		"Binding Sapling Root System",
		"Breeding Fauna",
		"Building Data Trees",
		"Bureacritizing Bureaucracies",
		"Calculating Inverse Probability Matrices",
		"Calculating Llama Expectoration Trajectory",
		"Calibrating Blue Skies",
		"Charging Ozone Layer",
		"Coalescing Cloud Formations",
		"Cohorting Exemplars",
		"Collecting Meteor Particles",
		"Compounding Inert Tessellations",
		"Compressing Fish Files",
		"Computing Optimal Bin Packing",
		"Concatenating Sub-Contractors",
		"Containing Existential Buffer",
		"Debarking Ark Ramp",
		"Debunching Unionized Commercial Services",
		"Deciding What Message to Display Next",
		"Decomposing Singular Values",
		"Decrementing Tectonic Plates",
		"Deleting Ferry Routes",
		"Depixelating Inner Mountain Surface Back Faces",
		"Depositing Slush Funds",
		"Destabilizing Economic Indicators",
		"Determining Width of Blast Fronts",
		"Deunionizing Bulldozers",
		"Dicing Models",
		"Diluting Livestock Nutrition Variables",
		"Downloading Satellite Terrain Data",
		"Exposing Flash Variables to Streak System",
		"Extracting Resources",
		"Factoring Pay Scale",
		"Fixing Election Outcome Matrix",
		"Flood-Filling Ground Water",
		"Flushing Pipe Network",
		"Gathering Particle Sources",
		"Generating Jobs",
		"Gesticulating Mimes",
		"Graphing Whale Migration",
		"Hiding Willio Webnet Mask",
		"Implementing Impeachment Routine",
		"Increasing Accuracy of RCI Simulators",
		"Increasing Magmafacation",
		"Initializing My Sim Tracking Mechanism",
		"Initializing Rhinoceros Breeding Timetable",
		"Initializing Robotic Click-Path AI",
		"Inserting Sublimated Messages",
		"Integrating Curves",
		"Integrating Illumination Form Factors",
		"Integrating Population Graphs",
		"Iterating Cellular Automata",
		"Lecturing Errant Subsystems",
		"Mixing Genetic Pool",
		"Modeling Object Components",
		"Mopping Occupant Leaks",
		"Normalizing Power",
		"Obfuscating Quigley Matrix",
		"Overconstraining Dirty Industry Calculations",
		"Partitioning City Grid Singularities",
		"Perturbing Matrices",
		"Pixalating Nude Patch",
		"Polishing Water Highlights",
		"Populating Lot Templates",
		"Preparing Sprites for Random Walks",
		"Prioritizing Landmarks",
		"Projecting Law Enforcement Pastry Intake",
		"Realigning Alternate Time Frames",
		"Reconfiguring User Mental Processes",
		"Relaxing Splines",
		"Removing Road Network Speed Bumps",
		"Removing Texture Gradients",
		"Removing Vehicle Avoidance Behavior",
		"Resolving GUID Conflict",
		"Reticulating Splines",
		"Retracting Phong Shader",
		"Retrieving from Back Store",
		"Reverse Engineering Image Consultant",
		"Routing Neural Network Infanstructure",
		"Scattering Rhino Food Sources",
		"Scrubbing Terrain",
		"Searching for Llamas",
		"Seeding Architecture Simulation Parameters",
		"Sequencing Particles",
		"Setting Advisor Moods",
		"Setting Inner Deity Indicators",
		"Setting Universal Physical Constants",
		"Sonically Enhancing Occupant-Free Timber",
		"Speculating Stock Market Indices",
		"Splatting Transforms",
		"Stratifying Ground Layers",
		"Sub-Sampling Water Data",
		"Synthesizing Gravity",
		"Synthesizing Wavelets",
		"Time-Compressing Simulator Clock",
		"Unable to Reveal Current Activity",
		"Weathering Buildings",
		"Zeroing Crime Network",
	}
)

var (
	box1   = `⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏`
	box2   = `⠋⠙⠚⠞⠖⠦⠴⠲⠳⠓`
	box3   = `⠄⠆⠇⠋⠙⠸⠰⠠⠰⠸⠙⠋⠇⠆`
	box4   = `⠋⠙⠚⠒⠂⠂⠒⠲⠴⠦⠖⠒⠐⠐⠒⠓⠋`
	box5   = `⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠴⠲⠒⠂⠂⠒⠚⠙⠉⠁`
	box6   = `⠈⠉⠋⠓⠒⠐⠐⠒⠖⠦⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈`
	box7   = `⠁⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈⠈`
	spin1  = `|/-\`
	spin2  = `◴◷◶◵`
	spin3  = `◰◳◲◱`
	spin4  = `◐◓◑◒`
	spin5  = `▉▊▋▌▍▎▏▎▍▌▋▊▉`
	spin6  = `▌▄▐▀`
	spin7  = `╫╪`
	spin8  = `■□▪▫`
	spin9  = `←↑→↓`
	spin10 = `⦾⦿`
	spin11 = `⌜⌝⌟⌞`
	spin12 = `┤┘┴└├┌┬┐`
	spin13 = `⇑⇗⇒⇘⇓⇙⇐⇖`
	spin14 = `☰☱☳☷☶☴`
	spin15 = `䷀䷪䷡䷊䷒䷗䷁䷖䷓䷋䷠䷫`
	def    = box1
	cursor = `■ ■ `

)

type spinner struct {
	frames []rune
	pos    int
	text   string
}

func (s *spinner) next() string {
	r := s.frames[s.pos%len(s.frames)]
	s.pos++
	return string(r)
}

func currentTimeString() string {
	t := time.Now()
	return fmt.Sprintf("[white::b]%s ", t.Format(root.TimeFormat))
}

func Refresh(app *tview.Application) {
	tick := time.NewTicker(refreshInterval)
	for {
		select {
		case <-tick.C:
			app.Draw()
		}
	}
}

func Main(a root.Args, s map[string]string, prim tview.Primitive) tview.Primitive {
	title := fmt.Sprintf("[yellow]%s", a.Module)
	info := fmt.Sprintf(" » [white::u]%s", s["info"])
	spin := &spinner{
		frames: []rune(cursor),
	}
	drawHeader := func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
		tview.Print(screen, fmt.Sprintf("Ready %s", spin.next()), x+1, height/2, width, tview.AlignLeft, tcell.ColorLime)
		tview.Print(screen, title, x-4, height/2, width, tview.AlignCenter, tcell.ColorLime)
		tview.Print(screen, currentTimeString(), x, height/2, width, tview.AlignRight, tcell.ColorLime)
		return 0, 0, 0, 0
	}
	drawFooter := func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
		tview.Print(screen, info, x+1, y, width, tview.AlignLeft, tcell.ColorGreen)
		//tview.Print(screen, message[rand.Intn(len(message))], x+1, y, width, tview.AlignLeft, tcell.ColorLime)
		//tview.Print(screen, title, x-1, y, width, tview.AlignCenter, tcell.ColorLime)
		tview.Print(screen, fmt.Sprintf("[cyan] %s@%s  ", a.Username, a.Hostname), x, y, width, tview.AlignRight, tcell.ColorLime)
		return 0, 0, 0, 0
	}
	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().SetDrawFunc(drawHeader), 0, 2, false).
		AddItem(tview.NewFrame(prim).SetBorders(2, 2, 2, 2, 4, 4), 0, 12, true).
		AddItem(tview.NewBox().SetDrawFunc(drawFooter), 0, 2, false)
}

func Activity(a root.Args, prim tview.Primitive) tview.Primitive {
	title := fmt.Sprintf("[yellow]%s", a.Module)
	spin := &spinner{
		frames: []rune(def),
	}
	drawHeader := func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
		tview.Print(screen, fmt.Sprintf("Please wait %s", spin.next()), x+1, height/2, width, tview.AlignLeft, tcell.ColorLime)
		tview.Print(screen, title, x-4, height/2, width, tview.AlignCenter, tcell.ColorLime)
		tview.Print(screen, currentTimeString(), x, height/2, width, tview.AlignRight, tcell.ColorLime)
		return 0, 0, 0, 0
	}
	drawFooter := func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
		tview.Print(screen, message[rand.Intn(len(message))], x+1, y, width, tview.AlignLeft, tcell.ColorLime)
		tview.Print(screen, fmt.Sprintf("[cyan] %s@%s  ", a.Username, a.Hostname), x, y, width, tview.AlignRight, tcell.ColorLime)
		return 0, 0, 0, 0
	}
	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().SetDrawFunc(drawHeader), 0, 2, false).
		AddItem(tview.NewFrame(prim).SetBorders(2, 2, 2, 2, 4, 4), 0, 12, true).
		AddItem(tview.NewBox().SetDrawFunc(drawFooter), 0, 2, false)
}

func End(a root.Args, text string, status string, code int) tview.Primitive {
	return tview.NewFrame(tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetText(text)).
		SetBorders(2, 2, 2, 2, 4, 4).
		AddText(a.Module, true, tview.AlignLeft, tcell.ColorYellow).
		AddText(a.Hostname, true, tview.AlignCenter, tcell.ColorWhite).
		AddText(time.Now().Format(root.TimeFormat), true, tview.AlignRight, tcell.ColorGreen).
		AddText(status, false, tview.AlignLeft, tcell.ColorDefault).
		AddText(fmt.Sprintf("%d", code), false, tview.AlignCenter, tcell.ColorRed).
		AddText(fmt.Sprintf("%s(%s)", a.Username, a.Uid), false, tview.AlignRight, tcell.ColorWhite)
}

func Back(text string, app *tview.Application, pages *tview.Pages) tview.Primitive {
	return tview.NewModal().
		SetText(text).
		AddButtons([]string{"Go back to Main screen", "Quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonIndex == 0 {
				pages.SwitchToPage("Main")
			} else {
				app.Stop()
			}
		})
}

func Center(width, height int, p tview.Primitive) tview.Primitive {
	return tview.NewFlex().
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(tview.NewBox(), 0, 1, false), width, 1, true).
		AddItem(tview.NewBox(), 0, 1, false)
}
