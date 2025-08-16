import { Q as copy_payload, R as assign_payload, C as pop, z as push, G as attr_class, T as clsx, F as attr, I as escape_html, J as stringify } from "../../../../chunks/index.js";
import { P as PlayerLookup } from "../../../../chunks/PlayerLookup.js";
function InputForm($$payload, $$props) {
  push();
  let formState = {
    current_thieving_level: 55,
    target_thieving_level: 99,
    has_ardy_med: false,
    has_thieving_cape: false,
    has_rogues_outfit: false,
    has_shadow_veil: false,
    hourly_pickpockets: 1e3,
    food_heal_amount: 20,
    food_cost: 100
  };
  let playerUsername = "";
  let playerLookupLoading = false;
  let isLoading = false;
  const inputBaseClasses = "block w-full rounded-md border-0 py-2 px-3.5 bg-gray-700/50 text-theme-text-primary shadow-sm ring-1 ring-inset ring-theme-border-input placeholder:text-theme-text-tertiary focus:ring-2 focus:ring-inset focus:ring-theme-accent sm:text-sm sm:leading-6 shadow-inner-border transition-colors duration-150";
  const fieldsetLegendClasses = "text-base font-semibold leading-7 text-theme-text-primary mb-1";
  const labelBaseClasses = "block text-xs font-medium text-theme-text-secondary mb-1";
  const checkboxClasses = "h-4 w-4 text-theme-accent focus:ring-theme-accent focus:ring-offset-theme-card-bg border-theme-border-input rounded";
  let $$settled = true;
  let $$inner_payload;
  function $$render_inner($$payload2) {
    $$payload2.out += `<form class="bg-theme-card-bg shadow-card rounded-lg border border-theme-border p-6 space-y-8"><h2 class="text-h3 text-theme-text-primary border-b border-theme-border pb-4">Ardougne Knight Inputs</h2> <fieldset class="space-y-4"><legend${attr_class(clsx(fieldsetLegendClasses))}>Player Lookup (Optional)</legend> <div>`;
    PlayerLookup($$payload2, {
      placeholder: "Enter OSRS username to auto-fill thieving level",
      buttonText: "Load Stats",
      get username() {
        return playerUsername;
      },
      set username($$value) {
        playerUsername = $$value;
        $$settled = false;
      },
      get loading() {
        return playerLookupLoading;
      },
      set loading($$value) {
        playerLookupLoading = $$value;
        $$settled = false;
      }
    });
    $$payload2.out += `<!----> <p class="text-theme-text-tertiary text-xs mt-1">Automatically loads your current thieving level from OSRS hiscores</p></div></fieldset> <fieldset class="space-y-4"><legend${attr_class(clsx(fieldsetLegendClasses))}>Character Stats</legend> <div class="grid grid-cols-2 gap-4"><div><label for="current-level"${attr_class(clsx(labelBaseClasses))}>Current Thieving Level:</label> <input type="number" id="current-level"${attr("value", formState.current_thieving_level)} min="1" max="98" step="1" placeholder="e.g., 55"${attr_class(`${stringify(inputBaseClasses)} mt-1`)} required/></div> <div><label for="target-level"${attr_class(clsx(labelBaseClasses))}>Target Thieving Level:</label> <input type="number" id="target-level"${attr("value", formState.target_thieving_level)} min="2" max="99" step="1" placeholder="e.g., 99"${attr_class(`${stringify(inputBaseClasses)} mt-1`)} required/></div></div></fieldset> <fieldset class="space-y-4"><legend${attr_class(clsx(fieldsetLegendClasses))}>Equipment &amp; Bonuses</legend> <div class="grid grid-cols-2 gap-4"><label class="flex items-center space-x-3"><input type="checkbox"${attr("checked", formState.has_ardy_med, true)}${attr_class(clsx(checkboxClasses))}/> <span class="text-sm text-theme-text-primary">Ardougne Medium Diary</span></label> <label class="flex items-center space-x-3"><input type="checkbox"${attr("checked", formState.has_thieving_cape, true)}${attr_class(clsx(checkboxClasses))}/> <span class="text-sm text-theme-text-primary">Thieving Cape</span></label> <label class="flex items-center space-x-3"><input type="checkbox"${attr("checked", formState.has_rogues_outfit, true)}${attr_class(clsx(checkboxClasses))}/> <span class="text-sm text-theme-text-primary">Full Rogue's Outfit</span></label> <label class="flex items-center space-x-3"><input type="checkbox"${attr("checked", formState.has_shadow_veil, true)}${attr_class(clsx(checkboxClasses))}/> <span class="text-sm text-theme-text-primary">Shadow Veil Spell</span></label></div></fieldset> <fieldset class="space-y-4"><legend${attr_class(clsx(fieldsetLegendClasses))}>Training Settings</legend> <div class="space-y-4"><div><label for="hourly-pickpockets"${attr_class(clsx(labelBaseClasses))}>Hourly Pickpockets:</label> <input type="number" id="hourly-pickpockets"${attr("value", formState.hourly_pickpockets)} min="100" max="5000" step="50" placeholder="e.g., 1000"${attr_class(`${stringify(inputBaseClasses)} mt-1`)} required/></div> <div class="grid grid-cols-2 gap-4"><div><label for="food-heal"${attr_class(clsx(labelBaseClasses))}>Food Heal Amount:</label> <input type="number" id="food-heal"${attr("value", formState.food_heal_amount)} min="1" max="99" step="1" placeholder="e.g., 20"${attr_class(`${stringify(inputBaseClasses)} mt-1`)} required/></div> <div><label for="food-cost"${attr_class(clsx(labelBaseClasses))}>Food Cost (GP):</label> <input type="number" id="food-cost"${attr("value", formState.food_cost)} min="1" max="10000" step="1" placeholder="e.g., 100"${attr_class(`${stringify(inputBaseClasses)} mt-1`)} required/></div></div></div></fieldset> <button type="submit"${attr("disabled", isLoading, true)} class="w-full flex justify-center py-3 px-4 border border-transparent rounded-lg shadow-button text-sm font-semibold text-white bg-theme-accent hover:bg-theme-accent-hover focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-theme-card-bg focus:ring-theme-accent disabled:opacity-60 disabled:cursor-not-allowed transition-colors duration-150">${escape_html("Calculate Ardougne Knights")}</button></form>`;
  }
  do {
    $$settled = true;
    $$inner_payload = copy_payload($$payload);
    $$render_inner($$inner_payload);
  } while (!$$settled);
  assign_payload($$payload, $$inner_payload);
  pop();
}
function _page($$payload, $$props) {
  push();
  const toolConfig = {
    name: "Ardougne Knight Calculator",
    description: "Calculate your Thieving XP, GP, and more.",
    iconSrc: "/images/tools/knight_of_ardougne.png"
  };
  $$payload.out += `<div class="max-w-6xl mx-auto"><div class="mb-6"><a href="/tools" class="inline-flex items-center text-sm font-medium text-theme-accent hover:text-theme-accent-hover transition-colors group"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-5 h-5 mr-1.5 transform transition-transform group-hover:-translate-x-1"><path fill-rule="evenodd" d="M17 10a.75.75 0 0 1-.75.75H5.56l2.72 2.72a.75.75 0 1 1-1.06 1.06l-4-4a.75.75 0 0 1 0-1.06l4-4a.75.75 0 0 1 1.06 1.06L5.56 9.25H16.25A.75.75 0 0 1 17 10Z" clip-rule="evenodd"></path></svg> Back to All Tools</a></div> <header class="mb-10 text-center"><div class="flex items-center justify-center mb-4"><div class="w-16 h-16 bg-gradient-to-r from-purple-500 to-indigo-600 rounded-xl flex items-center justify-center p-2 mr-4"><img${attr("src", toolConfig.iconSrc)}${attr("alt", `${stringify(toolConfig.name)} icon`)} class="w-full h-full object-contain"/></div> <h1 class="text-h1 text-theme-text-primary tracking-tight">${escape_html(toolConfig.name)}</h1></div> <p class="mt-3 text-lg text-theme-text-secondary">${escape_html(toolConfig.description)}</p></header> <div class="max-w-2xl mx-auto mb-8"><div class="bg-theme-bg-secondary rounded-card border border-theme-border-accent/20 p-6 shadow-card">`;
  InputForm($$payload);
  $$payload.out += `<!----></div></div> <div>`;
  {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--> `;
  {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--></div> <div class="max-w-4xl mx-auto mt-12"><div class="bg-theme-bg-secondary rounded-card border border-theme-border-accent/20 shadow-card overflow-hidden"><button class="w-full px-6 py-4 text-left bg-gradient-to-r from-purple-500/10 via-indigo-500/5 to-transparent hover:from-purple-500/15 hover:via-indigo-500/10 transition-all duration-200 flex items-center justify-between group"><div class="flex items-center space-x-3"><div class="w-8 h-8 bg-gradient-to-r from-purple-500 to-indigo-600 rounded-lg flex items-center justify-center flex-shrink-0"><svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg></div> <div><h3 class="text-lg font-semibold text-theme-text-primary">Calculation Methodology</h3> <p class="text-sm text-theme-text-secondary">Learn how Ardougne Knights XP rates and profit are calculated</p></div></div> <svg${attr_class(`w-5 h-5 text-theme-text-secondary transition-transform duration-200 ${stringify("")}`)} fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg></button> `;
  {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--></div></div></div>`;
  pop();
}
export {
  _page as default
};
