package osrs_api

type (
	MonsterMap map[int]*Monster

	Monster struct {
		Id          int
		Name        string
		CombatLevel int16
		Models      []int32

		Animations Animation
		Dimensions Dimension

		ItemDrops DropCollection

		WikiUrl string
		Valid   bool //valid wiki monster
	}

	Dimension struct {
		TileSpace int8
		ResizeX   int16
		ResizeY   int16
	}

	Animation struct {
		Stance        int32
		Walk          int32
		Rotate90Left  int32
		Rotate90Right int32
		Rotate180     int32
	}
)

func NewMonster(name string, url string) *Monster {
	return &Monster{
		Name:      name,
		WikiUrl:   url,
		ItemDrops: DropCollection{},
	}
}

func (m *Monster) Store() error {
	//TODO Persist

	return nil
}
