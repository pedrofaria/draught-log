<script>
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
        selectedItem = item;
        M.Sidenav.getInstance(document.querySelector('#attrview-sidenav')).open();
    }

    // Create WebSocket connection.
    const socket = new WebSocket('ws://' + location.host + '/api/stream');

    // Listen for messages
    socket.addEventListener('message', function (event) {
        let item = JSON.parse(event.data);
        addItem(item);
    });

    document.addEventListener('DOMContentLoaded', function() {
        var options = {
            edge: 'right',
            draggable: false,
            preventScrolling: true,
            onCloseEnd: function() {
                selectedItem = null;
            },
        }
        var elems = document.querySelectorAll('#attrview-sidenav');
        M.Sidenav.init(elems, options);
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

    #attrview-sidenav{
        width: 50%;
    }

    #attrview {
        padding: 1rem;
    }

    #attrview .msg {
        border: 1px solid #cccccc;
        background-color: #eeeeee;
        padding: 5px;
        overflow-wrap: break-word;
    }

    #attrview dd {
        overflow-wrap: break-word;
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
                    <a href="javascript:void" on:click={clearLogs}>Clear Logs</a>
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
                        <span class="level badge"
                              class:grey={item.level === "debug"}
                              class:blue={item.level === "info"}
                              class:yellow={item.level === "warn"}
                              class:red={item.level === "error"}
                        > </span>
                    </td>
                    <td>{formatDate(item.timestamp)}</td>
                    <td>{item.provider}</td>
                    <td class="field-message">{item.message}</td>
                </tr>
            {/each}
            </tbody>
        </table>

    </div>

    <div id="attrview-sidenav" class="sidenav z-depth-3">
        <div id="attrview">
            {#if selectedItem !== null}
                <div class="row">
                <span class="level"
                      class:grey={selectedItem.level === "debug"}
                      class:blue={selectedItem.level === "info"}
                      class:yellow={selectedItem.level === "warn"}
                      class:red={selectedItem.level === "error"}
                >{selectedItem.level}</span>

                    {selectedItem.timestamp}
                    <div class="divider"></div>
                </div>


                <div class="msg">{selectedItem.message}</div>

                {#if selectedItem.payload !== null}
                    <h6>Attributes</h6>
                    <dl>
                        {#each Object.keys(selectedItem.payload) as attr}
                            <dt class="blue-text text-darken-1">{attr}</dt>
                            <dd>{selectedItem.payload[attr]}</dd>
                        {/each}
                    </dl>
                {/if}
            {/if}
        </div>
    </div>
</div>
