<div class="
bg-white
rounded-3xl
shadow-md
overflow-hidden
hover:shadow-xl
transition
">


<img
src="{{.Image}}"
class="
w-full
h-60
object-cover
"
/>


<div class="p-6">


<h4 class="
text-xl
font-bold
">

{{.Name}}

</h4>


<div class="mt-4 text-gray-600 space-y-2">


<p>
📅 {{.Date}}
</p>


<p>
📍 {{.Venue}}
</p>


</div>



<a

href="/events/{{.ID}}"

class="
inline-block
mt-5
bg-blue-600
text-white
px-5
py-3
rounded-xl
hover:bg-blue-700
"

>

View Details →

</a>


</div>


</div>