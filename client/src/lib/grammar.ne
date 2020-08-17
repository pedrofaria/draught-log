@builtin "string.ne"
@builtin "whitespace.ne"

@{%
var expression = function(t, f, c) { return {type: t, field: f, content: c}; };
var content = function(t, c) { return {type: t, values: c} };
%}

main -> expression       {% (d) => [d[0]] %}
	| expression __ main {% (d) => [d[0], ...d[2]] %}

expression -> field_expression {% id %}
	| attr_expression          {% id %}

field_expression -> field ":" content {% (d) => expression("field",d[0],d[2]) %}

field -> [^@] [-a-zA-Z_]:+ {% (d) => d[0] + d[1].join("")%}

attr_expression -> attr ":" content {% (d) => expression("attribute",d[0],d[2]) %}

attr -> "@" [a-zA-Z0-9._]:+ {% (d) => d[1].join('') %}

content -> simple_content {% (d) => content("simple", d[0]) %}
	| parentesis_content {% id %}

simple_content -> word_content {% (d) => [d[0]] %}
	| dqstring                 {% (d) => [d[0]] %}
	| sqstring                 {% (d) => [d[0]] %}

word_content -> [-a-zA-Z0-9._]:+ {% (d) => d[0].join("") %}

parentesis_content -> "(" _ parentesis_expression _ ")" {% (d) => d[2] %}

parentesis_expression -> simple_content {% (d) => content("simple", d[0]) %}
	| parentesis_expression_OR {% (d) => content("expressionOR", d[0]) %}

parentesis_expression_OR -> simple_content __ "OR" __ parentesis_expression_OR {% (d) => d[0].concat(d[4]) %}
	| simple_content {% id %}
