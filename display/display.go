package display

type Display interface {
	Clean()
	Done()
	DrawAlive()
	DrawDead()
}
