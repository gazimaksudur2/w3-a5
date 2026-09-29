{{template "partials/header.tpl" .}}

<div class="
bg-white
rounded-3xl
shadow-xl
p-8
max-w-4xl
mx-auto
">

{{if .Error}}

<div class="
bg-red-100
text-red-700
p-4
rounded-xl
mb-6
">
{{.Error}}
</div>

{{else}}

{{if .Event.Image}}

<img
src="{{.Event.Image}}"
alt="{{.Event.Name}}"
class="
rounded-2xl
w-full
h-96
object-cover
"
/>

{{end}}


<h2 class="
text-4xl
font-bold
mt-8
">

{{.Event.Name}}

</h2>


<div class="
mt-5
text-gray-600
space-y-3
text-lg
">

<p>
📅 {{.Event.Date}}
</p>


<p>
📍 {{.Event.Venue}}
</p>

</div>


{{if .Event.Description}}

<p class="
mt-8
text-gray-700
leading-relaxed
">

{{.Event.Description}}

</p>

{{end}}



<a

href="{{.Event.TicketURL}}"

target="_blank"

class="
inline-block
mt-8
bg-green-600
hover:bg-green-700
text-white
px-8
py-4
rounded-xl
transition
"

>

View Tickets

</a>


{{end}}

</div>


{{template "partials/footer.tpl" .}}