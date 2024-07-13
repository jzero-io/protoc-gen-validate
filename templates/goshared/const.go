package goshared

const constTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.Const }}
		if {{ accessor . }} != {{ lit $r.GetConst }} {
			{{- if isEnum $f }}
			{{ if $r.GetConstMessage }}
				err := {{ err . $r.GetConstMessage }}
			{{ else }}
				err := {{ err . "value must equal " (enumVal $f $r.GetConst) }}
			{{ end }}
			{{- else }}
			{{ if $r.GetConstMessage }}
				err := {{ err . $r.GetConstMessage }}
			{{ else }}
				err := {{ err . "value must equal " $r.GetConst }}
			{{ end }}
			{{- end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`
