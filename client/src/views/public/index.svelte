<script>
    let MAX_ITEMS = 300;
    let items = [];
    let selectedItem = null;
    let totalLogs = 0;

    const formatter = new Intl.DateTimeFormat('en', {
        day: "2-digit",
        month: "short",
        hour12: false,
        hour: 'numeric',
        minute: '2-digit',
        second: 'numeric',
        fractionalSecondDigits: 1
    });

    function addItem(item) {
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
        items = items;
    }

    function selectItem(item) {
        console.log(item);
        selectedItem = item;
    }

    // Create WebSocket connection.
    const socket = new WebSocket('ws://localhost:5000/api/stream');

    // Listen for messages
    socket.addEventListener('message', function (event) {
        let item = JSON.parse(event.data);
        addItem(item);
    });

    function formatDate(ts) {
        const d = new Date(ts)
        return formatter.format(d)
    }


</script>

<style>
    #logpanel {
        font-size: 80%;
    }

    #logview {
        width: 100%;
        table-layout: fixed;
    }

    #logview td {
        padding: 0.3rem;
    }

    .col-level {
        width: 1.5rem;
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

    #attrview dd {
        overflow-wrap: break-word;
    }
</style>

<h1>Logs - {totalLogs}</h1>

<button on:click={clearLogs}>Clear</button>

<div id="logpanel" class="row">
    <div class="col s8">
        <table id="logview" class="highlight">
            <thead>
            <tr>
                <th class="col-level"></th>
                <th>Timestamp</th>
                <th>Provider</th>
                <th>Message</th>
            </tr>
            </thead>

            <tbody>
            {#each items as item (item.id)}
                <tr on:click={selectItem(item)}>
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

    {#if selectedItem !== null && selectedItem.payload !== null}
    <div id="attrview" class="col s3">
        <h5>Attributes</h5>

        <dl>
        {#each Object.keys(selectedItem.payload) as attr}
            <dt class="blue-text text-darken-1">{attr}</dt>
            <dd>{selectedItem.payload[attr]}</dd>
        {/each}
        </dl>
    </div>
    {/if}
</div>
