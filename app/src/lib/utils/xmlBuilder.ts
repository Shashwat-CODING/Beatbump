import { reduce } from "./collections";

export interface XMLNode {
	attrs?: Record<string, string | boolean | number>;
	name?: string;
	nodes?: Array<XMLNode>;
	text?: string;
	type: "element" | "declaration";
}

export interface XMLRoot {
	declaration: XMLNode;
	nodes?: Array<XMLNode>;
}
function buildAttrs(attributes: XMLNode["attrs"]) {
	let output = "";
	for (const attr in attributes) {
		output += ` ${attr}="${`${attributes[attr]}`
			.toString()
			.replace('"', '\\"')}"`;
	}
	return output;
}

function buildElement(node: XMLNode): string {
	let element = "<";

	if (node.name) element += node.name;
	if (node.attrs) element += buildAttrs(node.attrs);
	element += node.nodes || node.text ? ">" : "/>";
	if (node.text)
		element += `${node.text}`
			.replace(/&amp;/g, "&")
			.replace(/&/g, "&amp;")
			.replace(/</g, "&lt;")
			.replace(/>/g, "&gt;"); // desanitize to avoid double sanitization;
	if (node.nodes) element += buildElements(node.nodes);
	if (node.nodes || node.text) element += `</${node.name}>`;
	return element;
}
function buildDeclaration(node: XMLNode, depth = 0) {
	return `<?xml${buildAttrs(node["attrs"])}?>`;
}
function buildElements(nodes: XMLNode[]) {
	const elements = reduce(
		nodes,
		(output, node) => {
			switch (node.type) {
				case "element":
					return output + buildElement(node);
				case "declaration":
					return output + buildDeclaration(node);
			}
		},
		"",
	);
	return elements;
}

export function buildXML(input: XMLRoot) {
	let output = "";
	const depth = 0;
	if (input["declaration"]) {
		output += buildDeclaration(input["declaration"], depth);
	}
	if (input["nodes"]) {
		output += buildElements(input["nodes"]);
	}
	return output;
}
