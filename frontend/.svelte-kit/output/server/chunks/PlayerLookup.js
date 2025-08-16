import { z as push, V as fallback, F as attr, I as escape_html, S as bind_props, C as pop } from "./index.js";
function PlayerLookup($$payload, $$props) {
  push();
  let username = fallback($$props["username"], "");
  let loading = fallback($$props["loading"], false);
  let placeholder = fallback($$props["placeholder"], "Enter OSRS username...");
  let buttonText = fallback($$props["buttonText"], "Lookup Player");
  $$payload.out += `<div class="player-lookup"><div class="flex gap-2"><input type="text"${attr("value", username)}${attr("placeholder", placeholder)}${attr("disabled", loading, true)} class="block w-full rounded-md border-0 py-2 px-3.5 bg-gray-700/50 text-theme-text-primary shadow-sm ring-1 ring-inset ring-theme-border-input placeholder:text-theme-text-tertiary focus:ring-2 focus:ring-inset focus:ring-theme-accent sm:text-sm sm:leading-6 shadow-inner-border transition-colors duration-150 flex-1" maxlength="12"/> <button${attr("disabled", loading || !username.trim(), true)} class="px-4 py-2 bg-theme-accent text-white rounded-md hover:bg-theme-accent-hover disabled:opacity-50 disabled:cursor-not-allowed whitespace-nowrap font-medium transition-colors duration-150">${escape_html(loading ? "Loading..." : buttonText)}</button></div> `;
  {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--></div>`;
  bind_props($$props, { username, loading, placeholder, buttonText });
  pop();
}
export {
  PlayerLookup as P
};
