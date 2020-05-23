<script>
    let MAX_ITEMS = 300;
    let items = [];
    let totalLogs = 0;

    const formatter = new Intl.DateTimeFormat('en', {
        day: "2-digit",
        month: "short",
        hour12: false,
        hour: 'numeric',
        minute: '2-digit',
        second: '2-digit',
        fractionalSecondDigits: 3
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
    #log-view {
        table-layout: fixed;
    }

    #log-view td {
        white-space: nowrap;
        overflow: hidden;
    }

    #log-view th.msg {
    }

    .level {
        min-width: 1rem;
    }
</style>

<h1>Logs - {totalLogs}</h1>

<button on:click={clearLogs}>Clear</button>

<div class="row">
    <div class="col s12">
        <table class="highlight" id="log-view">
            <thead>
            <tr>
                <th></th>
                <th>Timestamp</th>
                <th>Provider</th>
                <th class="msg">Message</th>
            </tr>
            </thead>

            <tbody>
            {#each items as item (item.id)}
                <tr>
                    <td>
                        <span class="badge level"
                            class:grey={item.level === "debug"}
                            class:blue={item.level === "info"}
                            class:yellow={item.level === "warn"}
                            class:red={item.level === "error"}
                        > </span>
                    </td>
                    <td>{formatDate(item.timestamp)}</td>
                    <td>{item.provider}</td>
                    <td>{item.message}</td>
                </tr>
            {/each}
            </tbody>
        </table>

    </div>
</div>

