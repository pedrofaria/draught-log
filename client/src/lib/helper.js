
export const getClassFromLevel = function (level) {
    switch (level) {
        case "debug": return "grey";
        case "info": return "blue";
        case "warn": return "yellow darken-2";
        case "error": return "red";
        default: return "black";
    }
}

export const copyToClipboard = function (text) {
    navigator.clipboard.writeText(text);
    document.toaster({
        html: 'text copied!',
        displayLength: 1000,
        classes: 'rounded'
    })
}
