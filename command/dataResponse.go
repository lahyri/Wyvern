package command

const (
	baseURL = "http://dnd5eapi.co/api/"
)

type Class struct {
	Name   string `json:"name,omitempty"`
	HitDie int    `json:"hit_die,omitempty"`
	URL    string `json:"url,omitempty"`

	ProficiencyChoices []Choice     `json:"proficiency_choices,omitempty"`
	Proficiences       []SimpleInfo `json:"proficiencies,omitempty"`
	SavingThrows       []SimpleInfo `json:"saving_throws,omitempty"`
	StartingEquipment  SimpleInfo   `json:"starting_equipment,omitempty"`
	ClassLevels        SimpleInfo   `json:"class_levels,omitempty"`
	Subclasses         []SimpleInfo `json:"subclasses,omitempty"`
	Spellcasting       SimpleInfo   `json:"spellcasting,omitempty"`
}

type Choice struct {
	From    []SimpleInfo `json:"from,omitempty"`
	Type    string       `json:"type,omitempty"`
	Ammount int          `json:"choose,omitempty"`
}

type SimpleInfo struct {
	URL   string `json:"url,omitempty"`
	Name  string `json:"name,omitempty"`
	Class string `json:"class,omitempty"`
}

type Spell struct {
	Name          string       `json:"name,omitempty"`
	Description   []string     `json:"desc,omitempty"`
	HigherLevel   []string     `json:"higher_level,omitempty"`
	Range         string       `json:"range,omitempty"`
	Components    []string     `json:"components,omitempty"`
	Material      string       `json:"material,omitempty"`
	Ritual        string       `json:"ritual,omitempty"`
	Duration      string       `json:"duration,omitempty"`
	Concentration string       `json:"concentration,omitempty"`
	CastingTime   string       `json:"casting_time,omitempty"`
	Level         int          `json:"level,omitempty"`
	School        SimpleInfo   `json:"school,omitempty"`
	Classes       []SimpleInfo `json:"classes,omitempty"`
}
