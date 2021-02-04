package models

type Config struct {
	Settings struct{
		TemplateAuthor string
		TemplateCite string
		TemplateContactInfo string
	}

	Todo struct{
		DaybreakAcc string
		DaybreakPwd string
	}

	Weather struct {
		Location string
	}

	Stocks struct{
		Code string
	}

	Components struct{
		Todo bool
		Weather bool
		Stock bool
	}
}
