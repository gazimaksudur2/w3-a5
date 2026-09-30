<div
    class="group bg-white rounded-3xl overflow-hidden border border-slate-100 shadow-md hover:shadow-2xl hover:-translate-y-2 transition duration-300"
>
    <div class="relative h-64 overflow-hidden bg-slate-200">
        {{if .Image}}

        <img
            src="{{.Image}}"
            alt="{{.Name}}"
            loading="lazy"
            class="w-full h-full object-cover group-hover:scale-110 transition duration-500"
        />

        {{else}}

        <div class="w-full h-full flex items-center justify-center text-5xl">🎟</div>

        {{end}}

        <div
            class="absolute top-4 left-4 px-4 py-2 rounded-full bg-white/90 backdrop-blur text-blue-700 text-sm font-semibold"
        >
            Event
        </div>
    </div>

    <div class="p-6">
        <h4 class="text-xl font-bold text-slate-900 line-clamp-2">{{.Name}}</h4>

        <div class="mt-5 space-y-3 text-slate-600 text-sm">
            <div class="flex gap-2">
                <span> 📅 </span>

                <span> {{.Date}} </span>
            </div>

            <div class="flex gap-2">
                <span> 📍 </span>

                <span> {{.Venue}} </span>
            </div>
        </div>

        <a
            href="/events/{{.ID}}"
            class="mt-6 block text-center py-3 rounded-xl bg-gradient-to-r from-blue-600 to-indigo-600 text-white font-semibold hover:shadow-lg transition"
        >
            View Details →
        </a>
    </div>
</div>
