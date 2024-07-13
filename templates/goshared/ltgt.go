package goshared

const ltgtTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.Lt }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLt $r.GetGt }}
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val >= {{ $r.Lt }} {
					{{ if $r.LtGtMessage }}
						err := {{ err . $r.GetLtGtMessage }}
					{{ else }}
						err := {{ err . "value must be inside range (" $r.GetGt ", " $r.GetLt ")" }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val <= {{ $r.Gt }} {
					{{ if $r.LtGtMessage }}
						err := {{ err . $r.GetLtGtMessage }}
					{{ else }}
						err := {{ err . "value must be outside range [" $r.GetLt ", " $r.GetGt "]" }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{  if gt $r.GetLt $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val >= {{ $r.Lt }} {
					{{ if $r.LtGteMessage }}
						err := {{ err . $r.GetLtGteMessage }}
					{{ else }}
						err := {{ err . "value must be inside range [" $r.GetGte ", " $r.GetLt ")" }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val < {{ $r.Gte }} {
					{{ if $r.LtGteMessage }}
						err := {{ err . $r.GetLtGteMessage }}
					{{ else }}
						err := {{ err . "value must be outside range [" $r.GetLt ", " $r.GetGte ")" }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} >= {{ $r.Lt }} {
				{{ if $r.LtMessage }}
					err := {{ err . $r.GetLtMessage }}
				{{ else }}
					err := {{ err . "value must be less than " $r.GetLt }}
				{{ end }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.Lte }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLte $r.GetGt }}
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val > {{ $r.Lte }} {
					{{ if $r.LteGtMessage }}
						err := {{ err . $r.GetLteGtMessage }}
					{{ else }}
						err := {{ err . "value must be inside range (" $r.GetGt ", " $r.GetLte "]" }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val <= {{ $r.Gt }} {
					{{ if $r.LteGtMessage }}
						err := {{ err . $r.GetLteGtMessage }}
					{{ else }}
						err := {{ err . "value must be outside range (" $r.GetLte ", " $r.GetGt "]" }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{ if gt $r.GetLte $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val > {{ $r.Lte }} {
					{{ if $r.LteGteMessage }}
						err := {{ err . $r.GetLteGteMessage }}
					{{ else }}
						err := {{ err . "value must be inside range [" $r.GetGte ", " $r.GetLte "]" }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val < {{ $r.Gte }} {
					{{ if $r.LteGteMessage }}
						err := {{ err . $r.GetLteGteMessage }}
					{{ else }}
						err := {{ err . "value must be outside range (" $r.GetLte ", " $r.GetGte ")" }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} > {{ $r.Lte }} {
				{{ if $r.LteMessage }}
					err := {{ err . $r.GetLteMessage }}
				{{ else }}
					err := {{ err . "value must be less than or equal to " $r.GetLte }}
				{{ end }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.Gt }}
		if {{ accessor . }} <= {{ $r.Gt }} {
			{{ if $r.GtMessage }}
				err := {{ err . $r.GetGtMessage }}
			{{ else }}
				err := {{ err . "value must be greater than " $r.GetGt }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.Gte }}
		if {{ accessor . }} < {{ $r.Gte }} {
			{{ if $r.GteMessage }}
				err := {{ err . $r.GetGteMessage }}
			{{ else }}
				err := {{ err . "value must be greater than or equal to " $r.GetGte }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`
