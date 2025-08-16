import { Q as copy_payload, R as assign_payload, C as pop, z as push, E as ensure_array_like, I as escape_html, F as attr, G as attr_class, J as stringify } from "../../../../chunks/index.js";
import { P as PlayerLookup } from "../../../../chunks/PlayerLookup.js";
function InputForm($$payload, $$props) {
  push();
  let formData = { currentLevel: 27, targetLevel: 77 };
  let playerUsername = "";
  let playerLookupLoading = false;
  let isLoading = false;
  let validationErrors = {};
  const levelPresets = [
    {
      label: "Level 27 → 77 (Early GOTR)",
      current: 27,
      target: 77
    },
    {
      label: "Level 77 → 99 (Optimal GOTR)",
      current: 77,
      target: 99
    },
    {
      label: "Level 50 → 99",
      current: 50,
      target: 99
    },
    {
      label: "Level 27 → 99 (Full GOTR)",
      current: 27,
      target: 99
    }
  ];
  let $$settled = true;
  let $$inner_payload;
  function $$render_inner($$payload2) {
    const each_array = ensure_array_like(levelPresets);
    $$payload2.out += `<form class="space-y-6"><div class="space-y-4"><div class="block text-sm font-semibold text-theme-text-primary mb-3">Player Lookup (Optional)</div> <div>`;
    PlayerLookup($$payload2, {
      placeholder: "Enter OSRS username to auto-fill current level",
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
    $$payload2.out += `<!----> <p class="text-theme-text-tertiary text-xs mt-1">Automatically loads your current runecrafting level from OSRS hiscores</p></div></div> <div class="mb-6"><div class="block text-sm font-semibold text-theme-text-primary mb-3">Quick Presets</div> <div class="grid grid-cols-1 sm:grid-cols-2 gap-3"><!--[-->`;
    for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
      let preset = each_array[$$index];
      $$payload2.out += `<button type="button" class="text-left p-3 bg-theme-bg-tertiary hover:bg-theme-bg-elevated border border-theme-border-subtle hover:border-theme-accent-primary/30 rounded-lg transition-colors duration-200 text-sm"><div class="font-medium text-theme-text-primary">${escape_html(preset.label)}</div> <div class="text-xs text-theme-text-tertiary mt-1">${escape_html(preset.current)} → ${escape_html(preset.target)}</div></button>`;
    }
    $$payload2.out += `<!--]--></div></div> <div class="grid grid-cols-1 sm:grid-cols-2 gap-6"><div class="space-y-2"><label for="currentLevel" class="block text-sm font-semibold text-theme-text-primary">Current Runecrafting Level</label> <div class="relative"><div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"><div class="w-5 h-5 bg-gradient-to-br from-purple-500 to-blue-600 rounded flex items-center justify-center"><span class="text-white text-xs font-bold">RC</span></div></div> <input type="number" id="currentLevel"${attr("value", formData.currentLevel)} min="27" max="126"${attr_class("w-full pl-12 pr-4 py-3 bg-theme-bg-elevated border border-theme-border-primary rounded-lg focus:ring-2 focus:ring-theme-accent-primary focus:border-theme-accent-primary transition-colors text-theme-text-primary placeholder-theme-text-tertiary", void 0, {
      "border-red-500": validationErrors.currentLevel
    })} placeholder="Enter current level"/></div> `;
    if (validationErrors.currentLevel) {
      $$payload2.out += "<!--[-->";
      $$payload2.out += `<p class="text-red-400 text-xs">${escape_html(validationErrors.currentLevel)}</p>`;
    } else {
      $$payload2.out += "<!--[!-->";
    }
    $$payload2.out += `<!--]--></div> <div class="space-y-2"><label for="targetLevel" class="block text-sm font-semibold text-theme-text-primary">Target Runecrafting Level</label> <div class="relative"><div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"><div class="w-5 h-5 bg-gradient-to-br from-green-500 to-emerald-600 rounded flex items-center justify-center"><span class="text-white text-xs font-bold">🎯</span></div></div> <input type="number" id="targetLevel"${attr("value", formData.targetLevel)} min="27" max="126"${attr_class("w-full pl-12 pr-4 py-3 bg-theme-bg-elevated border border-theme-border-primary rounded-lg focus:ring-2 focus:ring-theme-accent-primary focus:border-theme-accent-primary transition-colors text-theme-text-primary placeholder-theme-text-tertiary", void 0, { "border-red-500": validationErrors.targetLevel })} placeholder="Enter target level"/></div> `;
    if (validationErrors.targetLevel) {
      $$payload2.out += "<!--[-->";
      $$payload2.out += `<p class="text-red-400 text-xs">${escape_html(validationErrors.targetLevel)}</p>`;
    } else {
      $$payload2.out += "<!--[!-->";
    }
    $$payload2.out += `<!--]--></div></div> <div class="bg-gradient-to-r from-purple-500/10 via-purple-600/5 to-transparent border border-purple-500/20 rounded-lg p-4"><div class="flex items-start space-x-3"><div class="w-6 h-6 bg-purple-500 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5"><span class="text-white text-xs">ℹ️</span></div> <div><h4 class="text-sm font-semibold text-purple-400 mb-1">About GOTR</h4> <p class="text-xs text-theme-text-secondary leading-relaxed">Guardians of the Rift provides excellent Runecrafting XP rates and valuable rewards. 
                    This calculator estimates your progress, pet chances, and potential loot based on average game performance.</p></div></div></div> <button type="submit"${attr("disabled", isLoading, true)} class="w-full bg-theme-accent hover:bg-theme-accent-hover disabled:bg-theme-bg-tertiary disabled:cursor-not-allowed text-white font-semibold py-4 px-6 rounded-button shadow-button hover:shadow-button-hover transition-all duration-200 flex items-center justify-center">`;
    {
      $$payload2.out += "<!--[!-->";
      $$payload2.out += `<div class="flex items-center"><span class="mr-2">🧙‍♂️</span> Calculate GOTR Training</div>`;
    }
    $$payload2.out += `<!--]--></button></form>`;
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
    name: "Guardians of the Rift Calculator",
    description: "Calculate XP rates, training time, reward rolls, and pet chances for GOTR based on your current and target levels.",
    iconSrc: "/images/skills/runecraft.png"
  };
  $$payload.out += `<div class="max-w-6xl mx-auto"><div class="mb-6"><a href="/tools" class="inline-flex items-center text-sm font-medium text-theme-accent hover:text-theme-accent-hover transition-colors group"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-5 h-5 mr-1.5 transform transition-transform group-hover:-translate-x-1"><path fill-rule="evenodd" d="M17 10a.75.75 0 0 1-.75.75H5.56l2.72 2.72a.75.75 0 1 1-1.06 1.06l-4-4a.75.75 0 0 1 0-1.06l4-4a.75.75 0 0 1 1.06 1.06L5.56 9.25H16.25A.75.75 0 0 1 17 10Z" clip-rule="evenodd"></path></svg> Back to All Tools</a></div> <header class="mb-10 text-center"><div class="flex items-center justify-center mb-4"><div class="w-16 h-16 bg-gradient-to-r from-purple-500 to-blue-600 rounded-xl flex items-center justify-center p-2 mr-4"><img${attr("src", toolConfig.iconSrc)}${attr("alt", `${stringify(toolConfig.name)} icon`)} class="w-full h-full object-contain"/></div> <h1 class="text-h1 text-theme-text-primary tracking-tight">${escape_html(toolConfig.name)}</h1></div> <p class="mt-3 text-lg text-theme-text-secondary">${escape_html(toolConfig.description)}</p></header> <div class="max-w-2xl mx-auto mb-8"><div class="bg-theme-bg-secondary rounded-card border border-theme-border-accent/20 p-6 shadow-card">`;
  InputForm($$payload);
  $$payload.out += `<!----></div></div> <div>`;
  {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--> `;
  {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--></div> <div class="max-w-4xl mx-auto mt-12"><div class="bg-theme-bg-secondary rounded-card border border-theme-border-accent/20 shadow-card overflow-hidden"><button class="w-full px-6 py-4 text-left bg-gradient-to-r from-blue-500/10 via-purple-500/5 to-transparent hover:from-blue-500/15 hover:via-purple-500/10 transition-all duration-200 flex items-center justify-between group"><div class="flex items-center space-x-3"><div class="w-8 h-8 bg-gradient-to-r from-blue-500 to-purple-600 rounded-lg flex items-center justify-center flex-shrink-0"><svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg></div> <div><h3 class="text-lg font-semibold text-theme-text-primary">Calculation Methodology</h3> <p class="text-sm text-theme-text-secondary">Learn how GOTR XP rates and rewards are calculated</p></div></div> <svg${attr_class(`w-5 h-5 text-theme-text-secondary transition-transform duration-200 ${stringify("")}`)} fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg></button> `;
  {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--></div></div></div>`;
  pop();
}
export {
  _page as default
};
