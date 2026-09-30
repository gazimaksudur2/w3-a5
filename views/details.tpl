{{template "partials/header.tpl" .}} {{if .Error}}

<div class="max-w-3xl mx-auto bg-red-100 text-red-700 p-6 rounded-3xl font-medium">{{.Error}}</div>

{{else}}

<section class="max-w-5xl mx-auto">
    <!-- Hero Image -->

    <div class="relative rounded-3xl overflow-hidden shadow-2xl">
        <img src="{{.Event.Image}}" alt="{{.Event.Name}}" class="w-full h-[500px] object-cover" />

        <div class="absolute inset-0 bg-gradient-to-t from-black/70 via-black/20 to-transparent"></div>

        <div class="absolute bottom-8 left-8 right-8 text-white">
            <h1 class="text-4xl md:text-6xl font-extrabold leading-tight">{{.Event.Name}}</h1>
        </div>
    </div>

    <!-- Event Information -->

    <div class="mt-10 bg-white rounded-3xl shadow-xl border border-slate-100 p-8 md:p-10">
        <div class="grid md:grid-cols-2 gap-6">
            <div class="bg-blue-50 rounded-2xl p-5">
                <p class="text-sm text-blue-600 font-semibold uppercase tracking-wide">Date</p>

                <p class="mt-2 text-xl font-bold text-slate-800">📅 {{.Event.Date}}</p>
            </div>

            <div class="bg-indigo-50 rounded-2xl p-5">
                <p class="text-sm text-indigo-600 font-semibold uppercase tracking-wide">Venue</p>

                <p class="mt-2 text-xl font-bold text-slate-800">📍 {{.Event.Venue}}</p>
            </div>
        </div>

        {{if .Event.Description}}

        <div class="mt-10">
            <h2 class="text-2xl font-bold mb-4">About this event</h2>

            <p class="text-slate-600 leading-relaxed text-lg">{{.Event.Description}}</p>
        </div>

        {{end}}

        <div class="mt-10 flex flex-col sm:flex-row gap-4">
            <a
                href="/redirect/{{.Event.ID}}"
                target="_blank"
                class="flex-1 text-center py-4 rounded-2xl bg-gradient-to-r from-green-500 to-emerald-600 text-white font-bold text-lg shadow-lg hover:shadow-xl hover:scale-105 transition"
            >
                🎟 Buy Tickets
            </a>

            <a
                href="/"
                class="flex-1 text-center py-4 rounded-2xl bg-slate-100 text-slate-700 font-semibold hover:bg-slate-200 transition"
            >
                ← Explore More Events
            </a>
        </div>
    </div>
</section>

{{end}} {{template "partials/footer.tpl" .}}
