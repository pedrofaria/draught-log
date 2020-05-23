<script>
    import { Route } from 'svelte-router-spa'
    export let currentRoute
    const params = {}
</script>

<div class="app">
    <section class="section">
        <Route {currentRoute}  {params} />
    </section>
</div>