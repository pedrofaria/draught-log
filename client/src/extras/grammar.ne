@builtin "string.ne"
@builtin "whitespace.ne"

@{%
var expression = function(t, f, c) { return {type: t, field: f, content: c}; };
var content = function(t, c) { return {type: t, values: c} };
%}

main -> expression       {% (d) => [d[0]] %}
	| expression __ main {% (d) => [d[0], ...d[2]] %}

expression -> field_expression {% (d) => expression("field", d[0][0], d[0][2]) %}
	| attr_expression          {% (d) => expression("attribute", d[0][0], d[0][2]) %}

field_expression -> field ":" content

field -> _ field_names:+ {% (d) => d[1][0] %}
field_names -> "provider"i {% id %}
	| "level"i             {% id %}
	| "message"i           {% id %}

attr_expression -> attr ":" content

attr -> _ "@" [a-zA-Z0-9._]:+ {% (d) => d[2].join('') %}


content -> simple_content {% (d) => content("simple", d[0]) %}
	| parentesis_content {% id %}

simple_content -> word_content {% (d) => [d[0]] %}
	| dqstring                 {% (d) => [d[0]] %}
	| sqstring                 {% (d) => [d[0]] %}

word_content -> [-a-zA-Z0-9._]:+ {% (d) => d[0].join("") %}

parentesis_content -> "(" _ parentesis_expression _ ")" {% (d) => d[2] %}

parentesis_expression -> simple_content {% (d) => content("simple", d[0]) %}
	| parentesis_expression_OR {% (d) => content("expressionOR", d[0]) %}

parentesis_expression_OR -> simple_content __ "OR"i __ parentesis_expression_OR {% (d) => d[0].concat(d[4]) %}
	| simple_content {% id %}
