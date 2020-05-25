<script>
    import { Route } from 'svelte-router-spa'
    export let currentRoute
    const params = {}

    document.addEventListener('DOMContentLoaded', function() {
        var elems = document.querySelectorAll('.modal');
        M.Modal.init(elems, {});
    });
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

            <ul id="nav-mobile" class="right hide-on-med-and-down">
                <li><a href="" data-target="info-modal" class="modal-trigger"><i class="material-icons">info</i></a></li>
            </ul>
        </div>
    </nav>
</div>

<div id="info-modal" class="modal">
    <div class="modal-content">
        <h4>Informations</h4>
        <p>Draught Log is an opensource project made with love, beer and insomnia. Contribute!</p>
        <p>
            <a href="https://github.com/pedrofaria/draught-log" target="_blank">GitHub repository</a>
        </p>
    </div>
    <div class="modal-footer">
        <a href="#!" class="modal-close waves-effect waves-green btn-flat">close</a>
    </div>
</div>

<Route {currentRoute}  {params} />
