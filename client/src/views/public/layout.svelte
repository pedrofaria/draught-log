<script>
    import { Route } from 'svelte-router-spa'
    export let currentRoute
    const params = {}
</script>

<style>
    .brand-logo {
        margin-left: 10px;
    }
</style>

<div class="navbar-fixed">
    <nav>
        <div class="nav-wrapper amber darken-2">
            <div class="brand-logo center">Draught Log</div>

            <ul id="nav-mobile" class="left hide-on-med-and-down">
                <li><a href="">Something</a></li>
            </ul>

            <ul id="nav-mobile" class="right hide-on-med-and-down">
                <li><a href="">Something</a></li>
            </ul>
        </div>
    </nav>
</div>

<Route {currentRoute}  {params} />
