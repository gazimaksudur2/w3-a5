{{template "partials/header.tpl" .}}


<h2>
    Events in {{.City}}
</h2>


{{if .Error}}

<p class="error-message">
    {{.Error}}
</p>

{{end}}



<h3>
    Music Events
</h3>


{{range .Music}}

<div class="event-card">


    <h4>
        {{.Name}}
    </h4>


    {{if .Image}}

    <img 
        src="{{.Image}}"
        width="200"
        alt="{{.Name}}"
    >

    {{end}}



    <p>
        <strong>Date:</strong>
        {{.Date}}
    </p>


    <p>
        <strong>Venue:</strong>
        {{.Venue}}
    </p>



    <a href="/events/{{.ID}}">
        View Details
    </a>


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

<div class="event-card">


    <h4>
        {{.Name}}
    </h4>


    {{if .Image}}

    <img 
        src="{{.Image}}"
        width="200"
        alt="{{.Name}}"
    >

    {{end}}



    <p>
        <strong>Date:</strong>
        {{.Date}}
    </p>


    <p>
        <strong>Venue:</strong>
        {{.Venue}}
    </p>



    <a href="/events/{{.ID}}">
        View Details
    </a>


</div>


<hr>


{{else}}

<p>
    No sports events found
</p>


{{end}}



{{template "partials/footer.tpl" .}}