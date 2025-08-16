import { N as slot } from "../../../chunks/index.js";
function _layout($$payload, $$props) {
  $$payload.out += `<div class="container mx-auto px-4 sm:px-6 lg:px-8 py-8 md:py-12"><!---->`;
  slot($$payload, $$props, "default", {});
  $$payload.out += `<!----></div>`;
}
export {
  _layout as default
};
