{{template "partials/header.tpl" .}}


<div class="event-details">


{{if .Error}}

    <div class="error-message">

        <h2>Error</h2>

        <p>{{.Error}}</p>

        <a href="/">Go Home</a>

    </div>


{{else}}


    <h2>{{.Event.Name}}</h2>


    {{if .Event.Image}}

        <img 
            src="{{.Event.Image}}"
            alt="{{.Event.Name}}"
            class="event-image"
        >

    {{end}}



    <div class="event-info">


        <p>
            <strong>Date:</strong>
            {{.Event.Date}}
        </p>


        <p>
            <strong>Venue:</strong>
            {{.Event.Venue}}
        </p>


        {{if .Event.Description}}

        <p>
            <strong>Description:</strong>
        </p>

        <p>
            {{.Event.Description}}
        </p>

        {{else}}

        <p>
            No description available.
        </p>

        {{end}}


    </div>



    <div class="actions">


        <a href="/events">
            ← Back to Events
        </a>



        <a 
            href="/redirect/{{.Event.ID}}"
            class="ticket-button"
        >
            View Tickets
        </a>


    </div>


{{end}}


</div>


{{template "partials/footer.tpl" .}}