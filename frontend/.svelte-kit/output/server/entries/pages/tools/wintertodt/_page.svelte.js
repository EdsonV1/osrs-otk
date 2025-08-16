import { Q as copy_payload, R as assign_payload, C as pop, z as push, G as attr_class, F as attr, T as clsx, U as maybe_selected, I as escape_html, J as stringify } from "../../../../chunks/index.js";
import { P as PlayerLookup } from "../../../../chunks/PlayerLookup.js";
function InputForm($$payload, $$props) {
  push();
  let formState = {
    currentLevel: 50,
    targetLevel: 99,
    strategy: "large_group",
    skillLevels: {
      herblore: 1,
      mining: 1,
      fishing: 1,
      crafting: 1,
      farming: 1,
      woodcutting: 1
    }
  };
  let playerUsername = "";
  let useLivePrices = false;
  let playerLookupLoading = false;
  let useCustomStrategy = false;
  let isLoading = false;
  const inputBaseClasses = "block w-full rounded-md border-0 py-2 px-3.5 bg-gray-700/50 text-theme-text-primary shadow-sm ring-1 ring-inset ring-theme-border-input placeholder:text-theme-text-tertiary focus:ring-2 focus:ring-inset focus:ring-theme-accent sm:text-sm sm:leading-6 shadow-inner-border transition-colors duration-150";
  const fieldsetLegendClasses = "text-base font-semibold leading-7 text-theme-text-primary mb-1";
  const labelBaseClasses = "block text-xs font-medium text-theme-text-secondary mb-1";
  let $$settled = true;
  let $$inner_payload;
  function $$render_inner($$payload2) {
    $$payload2.out += `<form class="bg-theme-card-bg shadow-card rounded-lg border border-theme-border p-6 space-y-8"><h2 class="text-h3 text-theme-text-primary border-b border-theme-border pb-4">Wintertodt Calculator</h2> <fieldset class="space-y-4"><legend${attr_class(clsx(fieldsetLegendClasses))}>Firemaking Levels</legend> <div class="grid grid-cols-1 md:grid-cols-2 gap-4"><div><label for="current-level"${attr_class(clsx(labelBaseClasses))}>Current Level (minimum 50):</label> <input type="number" id="current-level"${attr("value", formState.currentLevel)} min="50" max="99" step="1" placeholder="e.g., 75"${attr_class(`${stringify(inputBaseClasses)} mt-1`)} required/></div> <div><label for="target-level"${attr_class(clsx(labelBaseClasses))}>Target Level:</label> <input type="number" id="target-level"${attr("value", formState.targetLevel)} min="50" max="99" step="1" placeholder="e.g., 99"${attr_class(`${stringify(inputBaseClasses)} mt-1`)} required/></div></div></fieldset> <fieldset class="space-y-4"><legend${attr_class(clsx(fieldsetLegendClasses))}>Player Lookup (Optional)</legend> <div class="space-y-3"><div>`;
    PlayerLookup($$payload2, {
      placeholder: "Enter OSRS username to auto-fill stats",
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
    $$payload2.out += `<!----> <p class="text-theme-text-tertiary text-xs mt-1">Automatically loads your current skill levels from OSRS hiscores</p></div> <div class="flex items-center"><input type="checkbox" id="use-live-prices"${attr("checked", useLivePrices, true)} class="h-4 w-4 rounded border-theme-border bg-gray-700/50 text-theme-accent focus:ring-theme-accent focus:ring-2 focus:ring-offset-0"/> <label for="use-live-prices" class="ml-2 text-sm text-theme-text-secondary">Use live Grand Exchange prices (updates daily)</label></div></div></fieldset> <fieldset class="space-y-4"><legend${attr_class(clsx(fieldsetLegendClasses))}>Strategy</legend> <div><label for="strategy"${attr_class(clsx(labelBaseClasses))}>Strategy:</label> <select id="strategy"${attr_class(`${stringify(inputBaseClasses)} mt-1`)} required>`;
    $$payload2.select_value = formState.strategy;
    $$payload2.out += `<option value="large_group"${maybe_selected($$payload2, "large_group")}>Large Group (4 min/round, 750 points)</option><option value="solo"${maybe_selected($$payload2, "solo")}>Solo (15 min/round, 1000 points)</option><option value="efficient"${maybe_selected($$payload2, "efficient")}>Efficient Team (3.5 min/round, 600 points)</option>`;
    $$payload2.select_value = void 0;
    $$payload2.out += `</select></div> <div class="flex items-center space-x-2"><input type="checkbox" id="use-custom"${attr("checked", useCustomStrategy, true)} class="w-4 h-4 text-theme-accent bg-gray-700 border-gray-600 rounded focus:ring-theme-accent"/> <label for="use-custom" class="text-sm text-theme-text-secondary">Use custom points/time per round</label></div> `;
    {
      $$payload2.out += "<!--[!-->";
    }
    $$payload2.out += `<!--]--></fieldset> <fieldset class="space-y-4"><legend${attr_class(clsx(fieldsetLegendClasses))}>Other Skill Levels (for drops)</legend> <p class="text-xs text-theme-text-tertiary mb-4">Higher skill levels improve loot quality and quantity from supply crates</p> <div class="grid grid-cols-2 md:grid-cols-3 gap-4"><div><label for="herblore"${attr_class(clsx(labelBaseClasses))}>Herblore:</label> <input type="number" id="herblore"${attr("value", formState.skillLevels.herblore)} min="1" max="99" step="1"${attr_class(`${stringify(inputBaseClasses)} mt-1`)}/></div> <div><label for="mining"${attr_class(clsx(labelBaseClasses))}>Mining:</label> <input type="number" id="mining"${attr("value", formState.skillLevels.mining)} min="1" max="99" step="1"${attr_class(`${stringify(inputBaseClasses)} mt-1`)}/></div> <div><label for="fishing"${attr_class(clsx(labelBaseClasses))}>Fishing:</label> <input type="number" id="fishing"${attr("value", formState.skillLevels.fishing)} min="1" max="99" step="1"${attr_class(`${stringify(inputBaseClasses)} mt-1`)}/></div> <div><label for="crafting"${attr_class(clsx(labelBaseClasses))}>Crafting:</label> <input type="number" id="crafting"${attr("value", formState.skillLevels.crafting)} min="1" max="99" step="1"${attr_class(`${stringify(inputBaseClasses)} mt-1`)}/></div> <div><label for="farming"${attr_class(clsx(labelBaseClasses))}>Farming:</label> <input type="number" id="farming"${attr("value", formState.skillLevels.farming)} min="1" max="99" step="1"${attr_class(`${stringify(inputBaseClasses)} mt-1`)}/></div> <div><label for="woodcutting"${attr_class(clsx(labelBaseClasses))}>Woodcutting:</label> <input type="number" id="woodcutting"${attr("value", formState.skillLevels.woodcutting)} min="1" max="99" step="1"${attr_class(`${stringify(inputBaseClasses)} mt-1`)}/></div></div></fieldset> <button type="submit"${attr("disabled", isLoading, true)} class="w-full flex justify-center py-3 px-4 border border-transparent rounded-lg shadow-button text-sm font-semibold text-white bg-theme-accent hover:bg-theme-accent-hover focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-theme-card-bg focus:ring-theme-accent disabled:opacity-60 disabled:cursor-not-allowed transition-colors duration-150">${escape_html("Calculate Wintertodt")}</button></form>`;
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
    name: "Wintertodt Calculator",
    description: "Calculate experience, loot, and Phoenix pet chances from Wintertodt based on your Firemaking level and planned rounds.",
    iconSrc: "/images/skills/firemaking.png"
  };
  $$payload.out += `<div class="max-w-6xl mx-auto"><div class="mb-6"><a href="/tools" class="inline-flex items-center text-sm font-medium text-theme-accent hover:text-theme-accent-hover transition-colors group"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-5 h-5 mr-1.5 transform transition-transform group-hover:-translate-x-1"><path fill-rule="evenodd" d="M17 10a.75.75 0 0 1-.75.75H5.56l2.72 2.72a.75.75 0 1 1-1.06 1.06l-4-4a.75.75 0 0 1 0-1.06l4-4a.75.75 0 0 1 1.06 1.06L5.56 9.25H16.25A.75.75 0 0 1 17 10Z" clip-rule="evenodd"></path></svg> Back to All Tools</a></div> <header class="mb-10 text-center"><div class="flex items-center justify-center mb-4"><div class="w-16 h-16 bg-gradient-to-r from-orange-500 to-red-600 rounded-xl flex items-center justify-center p-2 mr-4"><img${attr("src", toolConfig.iconSrc)}${attr("alt", `${stringify(toolConfig.name)} icon`)} class="w-full h-full object-contain"/></div> <h1 class="text-h1 text-theme-text-primary tracking-tight">${escape_html(toolConfig.name)}</h1></div> <p class="mt-3 text-lg text-theme-text-secondary">${escape_html(toolConfig.description)}</p></header> <div class="max-w-2xl mx-auto mb-8"><div class="bg-theme-bg-secondary rounded-card border border-theme-border-accent/20 p-6 shadow-card">`;
  InputForm($$payload);
  $$payload.out += `<!----></div></div> <div>`;
  {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--> `;
  {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--></div> <div class="max-w-4xl mx-auto mt-12"><div class="bg-theme-bg-secondary rounded-card border border-theme-border-accent/20 shadow-card overflow-hidden"><button class="w-full px-6 py-4 text-left bg-gradient-to-r from-orange-500/10 via-red-500/5 to-transparent hover:from-orange-500/15 hover:via-red-500/10 transition-all duration-200 flex items-center justify-between group"><div class="flex items-center space-x-3"><div class="w-8 h-8 bg-gradient-to-r from-orange-500 to-red-600 rounded-lg flex items-center justify-center flex-shrink-0"><svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg></div> <div><h3 class="text-lg font-semibold text-theme-text-primary">Calculation Methodology</h3> <p class="text-sm text-theme-text-secondary">Learn how Wintertodt XP rates and rewards are calculated</p></div></div> <svg${attr_class(`w-5 h-5 text-theme-text-secondary transition-transform duration-200 ${stringify("")}`)} fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg></button> `;
  {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--></div></div></div>`;
  pop();
}
export {
  _page as default
};
