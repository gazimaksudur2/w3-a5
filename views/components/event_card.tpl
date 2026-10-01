<div class="bg-white rounded-2xl shadow-md overflow-hidden hover:shadow-xl transition duration-300">

    <div>
        <img
            src="{{.Image}}"
            alt="{{.Name}}"
            class="w-full h-40 object-cover"
        >
    </div>


    <div class="p-4">

        <h3 class="text-lg font-bold text-slate-800 line-clamp-2 min-h-[56px]">
            {{.Name}}
        </h3>


        <div class="mt-3 space-y-2 text-sm text-slate-500">

            <p>
                📅 {{.Date}}
            </p>

            <p class="line-clamp-1">
                📍 {{.Venue}}
            </p>

        </div>


        {{if .Description}}

        <p class="mt-3 text-sm text-slate-600 line-clamp-2">
            {{.Description}}
        </p>

        {{end}}


        <a
            href="/events/{{.ID}}"
            class="block text-center mt-4 px-4 py-2 rounded-lg bg-blue-600 text-white text-sm font-semibold hover:bg-blue-700 transition"
        >
            View Details
        </a>

    </div>

</div>