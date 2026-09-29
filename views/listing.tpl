{{template "partials/header.tpl" .}}


<h2>
Events in {{.City}}
</h2>


{{if .Error}}

<p>
{{.Error}}
</p>

{{end}}



<h3>
Music Events
</h3>


{{range .Music}}

<div>

<h4>
{{.Name}}
</h4>


<img 
src="{{.Image}}"
width="200"
/>


<p>
Date: {{.Date}}
</p>


<p>
Venue: {{.Venue}}
</p>


</div>

<hr>


{{else}}

<p>
No music events found
</p>


{{end}}



<h3>
Sports Events
</h3>


{{range .Sports}}

<div>

<h4>
{{.Name}}
</h4>


<img 
src="{{.Image}}"
width="200"
/>


<p>
Date: {{.Date}}
</p>


<p>
Venue: {{.Venue}}
</p>


</div>


<hr>


{{else}}

<p>
No sports events found
</p>


{{end}}



{{template "partials/footer.tpl" .}}