{{template "partials/header.tpl" .}}

{{if .Error}}

<div class="max-w-3xl mx-auto bg-red-50 border border-red-200 rounded-3xl p-8 text-center">
    <div class="text-5xl mb-4">⚠️</div>

    <h2 class="text-2xl font-bold text-red-700">Something went wrong</h2>

    <p class="mt-3 text-red-600">{{.Error}}</p>

    <a href="/" class="inline-block mt-6 px-6 py-3 rounded-xl bg-red-600 text-white font-semibold"> Try Again </a>
</div>

{{end}}

<section class="m-10 px-20">
    <div class="flex flex-col md:flex-row md:items-end justify-between gap-4">
        <div>
            <p class="text-blue-600 font-semibold uppercase tracking-wide text-sm">Events</p>

            <h2 class="text-4xl md:text-5xl font-extrabold mt-2">Events in {{.City}}</h2>

            <p class="mt-3 text-slate-500 text-lg">Discover concerts, sports and experiences happening nearby.</p>
        </div>

        <div class="px-5 py-3 rounded-full bg-blue-50 text-blue-700 font-semibold">📍 {{.City}}</div>
    </div>
</section>

{{if .Error}}

<div class="bg-red-100 text-red-700 p-20 rounded-2xl mb-10">{{.Error}}</div>

{{end}}

<!-- MUSIC EVENTS -->

<section class="lg:m-10 lg:px-20">
    <div class="flex items-center gap-3 mb-8">
        <div class="text-3xl">🎵</div>

        <h3 class="text-3xl font-bold">Music Events</h3>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
        {{range .Music}} {{template "components/event_card.tpl" .}} {{else}}

        <div class="col-span-full bg-slate-100 rounded-3xl p-10 text-center text-slate-500">No music events found.</div>

        {{end}}
    </div>
</section>

<!-- SPORTS EVENTS -->

<section class="mb-16 lg:m-10 lg:px-20">
    <div class="flex items-center gap-3 mb-8">
        <div class="text-3xl">🏟️</div>

        <h3 class="text-3xl font-bold">Sports Events</h3>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
        {{range .Sports}} {{template "components/event_card.tpl" .}} {{else}}

        <div class="col-span-full bg-slate-100 rounded-3xl p-10 text-center text-slate-500">
            No sports events found.
        </div>

        {{end}}
    </div>
</section>

{{template "partials/footer.tpl" .}}
