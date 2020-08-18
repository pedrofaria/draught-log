// Generated automatically by nearley, version 2.19.5
// http://github.com/Hardmath123/nearley
(function () {
function id(x) { return x[0]; }

var expression = function(t, f, c) { return {type: t, field: f, content: c}; };
var content = function(t, c) { return {type: t, values: c} };
var grammar = {
    Lexer: undefined,
    ParserRules: [
    {"name": "dqstring$ebnf$1", "symbols": []},
    {"name": "dqstring$ebnf$1", "symbols": ["dqstring$ebnf$1", "dstrchar"], "postprocess": function arrpush(d) {return d[0].concat([d[1]]);}},
    {"name": "dqstring", "symbols": [{"literal":"\""}, "dqstring$ebnf$1", {"literal":"\""}], "postprocess": function(d) {return d[1].join(""); }},
    {"name": "sqstring$ebnf$1", "symbols": []},
    {"name": "sqstring$ebnf$1", "symbols": ["sqstring$ebnf$1", "sstrchar"], "postprocess": function arrpush(d) {return d[0].concat([d[1]]);}},
    {"name": "sqstring", "symbols": [{"literal":"'"}, "sqstring$ebnf$1", {"literal":"'"}], "postprocess": function(d) {return d[1].join(""); }},
    {"name": "btstring$ebnf$1", "symbols": []},
    {"name": "btstring$ebnf$1", "symbols": ["btstring$ebnf$1", /[^`]/], "postprocess": function arrpush(d) {return d[0].concat([d[1]]);}},
    {"name": "btstring", "symbols": [{"literal":"`"}, "btstring$ebnf$1", {"literal":"`"}], "postprocess": function(d) {return d[1].join(""); }},
    {"name": "dstrchar", "symbols": [/[^\\"\n]/], "postprocess": id},
    {"name": "dstrchar", "symbols": [{"literal":"\\"}, "strescape"], "postprocess": 
        function(d) {
            return JSON.parse("\""+d.join("")+"\"");
        }
        },
    {"name": "sstrchar", "symbols": [/[^\\'\n]/], "postprocess": id},
    {"name": "sstrchar", "symbols": [{"literal":"\\"}, "strescape"], "postprocess": function(d) { return JSON.parse("\""+d.join("")+"\""); }},
    {"name": "sstrchar$string$1", "symbols": [{"literal":"\\"}, {"literal":"'"}], "postprocess": function joiner(d) {return d.join('');}},
    {"name": "sstrchar", "symbols": ["sstrchar$string$1"], "postprocess": function(d) {return "'"; }},
    {"name": "strescape", "symbols": [/["\\/bfnrt]/], "postprocess": id},
    {"name": "strescape", "symbols": [{"literal":"u"}, /[a-fA-F0-9]/, /[a-fA-F0-9]/, /[a-fA-F0-9]/, /[a-fA-F0-9]/], "postprocess": 
        function(d) {
            return d.join("");
        }
        },
    {"name": "_$ebnf$1", "symbols": []},
    {"name": "_$ebnf$1", "symbols": ["_$ebnf$1", "wschar"], "postprocess": function arrpush(d) {return d[0].concat([d[1]]);}},
    {"name": "_", "symbols": ["_$ebnf$1"], "postprocess": function(d) {return null;}},
    {"name": "__$ebnf$1", "symbols": ["wschar"]},
    {"name": "__$ebnf$1", "symbols": ["__$ebnf$1", "wschar"], "postprocess": function arrpush(d) {return d[0].concat([d[1]]);}},
    {"name": "__", "symbols": ["__$ebnf$1"], "postprocess": function(d) {return null;}},
    {"name": "wschar", "symbols": [/[ \t\n\v\f]/], "postprocess": id},
    {"name": "main", "symbols": ["expression"], "postprocess": (d) => [d[0]]},
    {"name": "main", "symbols": ["expression", "__", "main"], "postprocess": (d) => [d[0], ...d[2]]},
    {"name": "expression", "symbols": ["field_expression"], "postprocess": (d) => expression("field", d[0][0], d[0][2])},
    {"name": "expression", "symbols": ["attr_expression"], "postprocess": (d) => expression("attribute", d[0][0], d[0][2])},
    {"name": "field_expression", "symbols": ["field", {"literal":":"}, "content"]},
    {"name": "field$ebnf$1", "symbols": ["field_names"]},
    {"name": "field$ebnf$1", "symbols": ["field$ebnf$1", "field_names"], "postprocess": function arrpush(d) {return d[0].concat([d[1]]);}},
    {"name": "field", "symbols": ["_", "field$ebnf$1"], "postprocess": (d) => d[1][0]},
    {"name": "field_names$subexpression$1", "symbols": [/[pP]/, /[rR]/, /[oO]/, /[vV]/, /[iI]/, /[dD]/, /[eE]/, /[rR]/], "postprocess": function(d) {return d.join(""); }},
    {"name": "field_names", "symbols": ["field_names$subexpression$1"], "postprocess": id},
    {"name": "field_names$subexpression$2", "symbols": [/[lL]/, /[eE]/, /[vV]/, /[eE]/, /[lL]/], "postprocess": function(d) {return d.join(""); }},
    {"name": "field_names", "symbols": ["field_names$subexpression$2"], "postprocess": id},
    {"name": "field_names$subexpression$3", "symbols": [/[mM]/, /[eE]/, /[sS]/, /[sS]/, /[aA]/, /[gG]/, /[eE]/], "postprocess": function(d) {return d.join(""); }},
    {"name": "field_names", "symbols": ["field_names$subexpression$3"], "postprocess": id},
    {"name": "attr_expression", "symbols": ["attr", {"literal":":"}, "content"]},
    {"name": "attr$ebnf$1", "symbols": [/[a-zA-Z0-9._]/]},
    {"name": "attr$ebnf$1", "symbols": ["attr$ebnf$1", /[a-zA-Z0-9._]/], "postprocess": function arrpush(d) {return d[0].concat([d[1]]);}},
    {"name": "attr", "symbols": ["_", {"literal":"@"}, "attr$ebnf$1"], "postprocess": (d) => d[2].join('')},
    {"name": "content", "symbols": ["simple_content"], "postprocess": (d) => content("simple", d[0])},
    {"name": "content", "symbols": ["parentesis_content"], "postprocess": id},
    {"name": "simple_content", "symbols": ["word_content"], "postprocess": (d) => [d[0]]},
    {"name": "simple_content", "symbols": ["dqstring"], "postprocess": (d) => [d[0]]},
    {"name": "simple_content", "symbols": ["sqstring"], "postprocess": (d) => [d[0]]},
    {"name": "word_content$ebnf$1", "symbols": [/[-a-zA-Z0-9._]/]},
    {"name": "word_content$ebnf$1", "symbols": ["word_content$ebnf$1", /[-a-zA-Z0-9._]/], "postprocess": function arrpush(d) {return d[0].concat([d[1]]);}},
    {"name": "word_content", "symbols": ["word_content$ebnf$1"], "postprocess": (d) => d[0].join("")},
    {"name": "parentesis_content", "symbols": [{"literal":"("}, "_", "parentesis_expression", "_", {"literal":")"}], "postprocess": (d) => d[2]},
    {"name": "parentesis_expression", "symbols": ["simple_content"], "postprocess": (d) => content("simple", d[0])},
    {"name": "parentesis_expression", "symbols": ["parentesis_expression_OR"], "postprocess": (d) => content("expressionOR", d[0])},
    {"name": "parentesis_expression_OR$subexpression$1", "symbols": [/[oO]/, /[rR]/], "postprocess": function(d) {return d.join(""); }},
    {"name": "parentesis_expression_OR", "symbols": ["simple_content", "__", "parentesis_expression_OR$subexpression$1", "__", "parentesis_expression_OR"], "postprocess": (d) => d[0].concat(d[4])},
    {"name": "parentesis_expression_OR", "symbols": ["simple_content"], "postprocess": id}
]
  , ParserStart: "main"
}
if (typeof module !== 'undefined'&& typeof module.exports !== 'undefined') {
   module.exports = grammar;
} else {
   window.grammar = grammar;
}
})();
