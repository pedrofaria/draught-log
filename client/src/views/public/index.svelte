<script>
    let MAX_ITEMS = 300;
    let items = [];
    let totalLogs = 0;

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
</script>

<h1>Logs - {totalLogs}</h1>

<button on:click={clearLogs}>Clear</button>

<ul>
    {#each items as item (item.id)}
        <li>{item.message}</li>
    {/each}
</ul>
