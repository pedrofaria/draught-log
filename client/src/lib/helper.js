
export const getClassFromLevel = function (level) {
    switch (level) {
        case "debug": return "grey";
        case "info": return "blue";
        case "warn": return "yellow darken-2";
        case "error": return "red";
        default: return "black";
    }
}