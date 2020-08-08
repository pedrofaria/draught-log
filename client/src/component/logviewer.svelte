<script>
    import { fly } from 'svelte/transition';

    let MAX_ITEMS = 300;
    let items = [];
    let selectedItem = null;
    let totalLogs = 0;
    let searchTerm = "";


    function getPreparedItems(items, searchTerm) {
        // filter
        let list = items.filter(item => item.message.toLowerCase().indexOf(searchTerm.toLowerCase()) !== -1);

        // sort
        const sortField = 'timestamp';
        const sortAscending = false;

        function sortDirection(a, b, asc) {
            return asc ? a > b : b > a;
        }

        function comp(a, b) {
            if (sortDirection(a[sortField], b[sortField], sortAscending)) { return 1; }
            if (sortDirection(b[sortField], a[sortField], sortAscending)) { return -1; }
            return 0;
        }
        return list.sort(comp);
    }

    $: filteredList = getPreparedItems(items, searchTerm);

    const formatter = new Intl.DateTimeFormat('en', {
        day: "2-digit",
        month: "short",
        hour12: false,
        hour: 'numeric',
        minute: '2-digit',
        second: 'numeric',
        fractionalSecondDigits: 1
    });

    function formatDate(ts) {
        return formatter.format(ts);
    }

    function addItem(item) {
        item.timestamp = new Date(item.timestamp);
        items.unshift(item);
        totalLogs += 1;
        if (items.length > MAX_ITEMS) {
            items.splice(MAX_ITEMS - 1, items.length - MAX_ITEMS);
        }

        items = items;
    }

    function clearLogs() {
        items = [];
        totalLogs = 0;
    }

    function selectItem(item) {
        if (selectedItem !== null && item.id == selectedItem.id) {
            selectedItem = null;
            return;
        }

        selectedItem = item;
    }

    function getClassFromLevel(level) {
        switch (level) {
            case "debug": return "grey";
            case "info": return "blue";
            case "warn": return "yellow darken-2";
            case "error": return "red";
            default: return "black";
        }
    }

    // Create WebSocket connection.
    const socket = new WebSocket('ws://' + location.host + '/api/stream');

    // Listen for messages
    socket.addEventListener('message', function (event) {
        let item = JSON.parse(event.data);
        addItem(item);
    });

</script>

<style>
    #searchbox .input-field, #searchbox .input-field input {
        margin: 0;
    }

    #logpanel {
        font-size: 80%;
    }

    #logview {
        width: 100%;
        table-layout: fixed;
    }

    #logview tr.selected {
        background-color: #eed1b1;
    }

    #logview td {
        padding: 0.3rem;
    }

    .col-level {
        width: 1.5rem;
    }

    .col-timestamp {
        width: 11rem;
    }

    .col-provider {
        width: 12rem;
    }

    .field-message {
        white-space: nowrap;
        overflow: hidden;
    }

    .level {
        min-width: 1rem;
        max-width: 1rem;
        padding: 5px 0px;
    }

    #attrview {
        position: fixed;
        right: 0;
        top: 0;
        height: 100%;
        width: 50%;
        z-index: 1000;
        padding: 1rem;
        background-color: #ffffff;
        font-size: 80%;
        overflow: auto;
    }

    #attrview.debug {
        border-top: 5px solid #9e9e9e;
    }

    #attrview.info {
        border-top: 5px solid #2196F3;
    }

    #attrview.warn {
        border-top: 5px solid #fbc02d;
    }

    #attrview.error {
        border-top: 5px solid #F44336;
    }

    #attrview.critical {
        border-top: 5px solid #000000;
    }

    #attrview .msg {
        border: 1px solid #cccccc;
        padding: 5px;
        overflow-wrap: break-word;
    }

    #attrview dd {
        white-space: pre;
        overflow-wrap: break-word;
    }

    #attrview-header {
        height: 30px;
    }

    .attr-level {
        padding: 3px 5px;
        border-radius: 3px;
        font-weight: bold;
        color: #ffff;
    }
</style>

<div class="row">
    <div class="col s12">
        <div id="searchbox" class="card">
            <div class="card-content">
                <div class="input-field">
                    <input id="filterInput" type="text" class="validate" bind:value={searchTerm}>
                    <label for="filterInput">Filter</label>
                </div>
            </div>
            <div class="card-action">
                <div class="right">
                    <a href="#!" on:click={clearLogs}>Clear Logs</a>
                </div>
                <div>
                    Total logs: {totalLogs}
                </div>
            </div>
        </div>

    </div>
</div>

<div id="logpanel" class="row">
    <div class="col s12">
        <table id="logview" class="highlight">
            <thead>
            <tr>
                <th class="col-level"></th>
                <th class="col-timestamp">Timestamp</th>
                <th class="col-provider">Provider</th>
                <th class="col-message">Message</th>
            </tr>
            </thead>

            <tbody>
            {#each filteredList as item (item.id)}
                <tr on:click={selectItem(item)} class:selected={selectedItem == item}>
                    <td>
                        <span class="level badge {getClassFromLevel(item.level)}"> </span>
                    </td>
                    <td>{formatDate(item.timestamp)}</td>
                    <td>{item.provider}</td>
                    <td class="field-message">{item.message}</td>
                </tr>
            {/each}
            </tbody>
        </table>

    </div>
</div>

{#if selectedItem !== null}
<div id="attrview" class="z-depth-3 {selectedItem.level}" transition:fly="{{duration: 300, x: 500}}">
    <div class="row">
        <div id="attrview-header">
            <span class="attr-level {getClassFromLevel(selectedItem.level)}">{selectedItem.level}</span>

            {selectedItem.timestamp}
            <a href="#!"  class="right" on:click={selectItem(selectedItem)}>
                <i class="material-icons">close</i>
            </a>
        </div>

        <div class="divider"></div>
    </div>


    <div class="msg yellow lighten-5">{selectedItem.message}</div>

    {#if selectedItem.payload !== null}
        <h6>Attributes</h6>
        <dl>
            {#each Object.keys(selectedItem.payload) as attr}
                <dt class="orange-text text-darken-1">{attr}</dt>
                <dd>{selectedItem.payload[attr]}</dd>
            {/each}
        </dl>
    {/if}
</div>
{/if}
