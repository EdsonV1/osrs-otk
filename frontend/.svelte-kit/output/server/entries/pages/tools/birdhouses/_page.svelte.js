import { E as ensure_array_like, G as attr_class, F as attr, I as escape_html, J as stringify, T as clsx, C as pop, z as push } from "../../../../chunks/index.js";
function InputForm($$payload, $$props) {
  push();
  const birdhouseTypes = [
    {
      value: "regular",
      label: "Regular",
      hunterLevel: 5,
      iconSrc: "/images/birdhouse/bird_house.png"
    },
    {
      value: "oak",
      label: "Oak",
      hunterLevel: 14,
      iconSrc: "/images/birdhouse/oak_bird_house.png"
    },
    {
      value: "willow",
      label: "Willow",
      hunterLevel: 24,
      iconSrc: "/images/birdhouse/willow_bird_house.png"
    },
    {
      value: "teak",
      label: "Teak",
      hunterLevel: 34,
      iconSrc: "/images/birdhouse/teak_bird_house.png"
    },
    {
      value: "maple",
      label: "Maple",
      hunterLevel: 44,
      iconSrc: "/images/birdhouse/maple_bird_house.png"
    },
    {
      value: "mahogany",
      label: "Mahogany",
      hunterLevel: 49,
      iconSrc: "/images/birdhouse/mahogany_bird_house.png"
    },
    {
      value: "yew",
      label: "Yew",
      hunterLevel: 59,
      iconSrc: "/images/birdhouse/yew_bird_house.png"
    },
    {
      value: "magic",
      label: "Magic",
      hunterLevel: 74,
      iconSrc: "/images/birdhouse/magic_bird_house.png"
    },
    {
      value: "redwood",
      label: "Redwood",
      hunterLevel: 89,
      iconSrc: "/images/birdhouse/redwood_bird_house.png"
    }
  ];
  let formState = {
    selectedLogType: birdhouseTypes[0].value,
    totalBirdhouses: 0
  };
  let isLoading = false;
  const inputBaseClasses = "block w-full rounded-md border-0 py-2 px-3.5 bg-gray-700/50 text-theme-text-primary shadow-sm ring-1 ring-inset ring-theme-border-input placeholder:text-theme-text-tertiary focus:ring-2 focus:ring-inset focus:ring-theme-accent sm:text-sm sm:leading-6 shadow-inner-border transition-colors duration-150";
  const fieldsetLegendClasses = "text-base font-semibold leading-7 text-theme-text-primary mb-1";
  const labelBaseClasses = "block text-xs font-medium text-theme-text-secondary mb-1";
  const radioLabelClasses = "flex items-center space-x-3 p-3 border border-theme-border-subtle rounded-lg hover:border-theme-accent/70 cursor-pointer transition-all duration-150 ease-in-out";
  const radioInputClasses = "h-4 w-4 text-theme-accent focus:ring-theme-accent focus:ring-offset-theme-card-bg border-theme-border-input";
  const each_array = ensure_array_like(birdhouseTypes);
  $$payload.out += `<form class="bg-theme-card-bg shadow-card rounded-lg border border-theme-border p-6 space-y-8"><h2 class="text-h3 text-theme-text-primary border-b border-theme-border pb-4">Birdhouse Inputs</h2> <fieldset class="space-y-4"><legend${attr_class(`${stringify(fieldsetLegendClasses)} mb-3`)}>Type of Birdhouses</legend> <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-2 gap-3"><!--[-->`;
  for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
    let typeOpt = each_array[$$index];
    $$payload.out += `<label${attr_class(`${stringify(radioLabelClasses)} ${stringify(formState.selectedLogType === typeOpt.value ? "ring-2 ring-theme-accent border-theme-accent shadow-md" : "bg-gray-700/30 hover:bg-gray-700/60")}`)}><input type="radio" name="birdhouseLogType"${attr("checked", formState.selectedLogType === typeOpt.value, true)}${attr("value", typeOpt.value)}${attr_class(`${stringify(radioInputClasses)} sr-only`)}/> <img${attr("src", typeOpt.iconSrc)}${attr("alt", `${stringify(typeOpt.label)} icon`)} class="w-8 h-8 object-contain flex-shrink-0"/> <span class="text-sm font-medium text-theme-text-primary">${escape_html(typeOpt.label)}</span></label>`;
  }
  $$payload.out += `<!--]--></div></fieldset> <fieldset class="space-y-4"><legend${attr_class(clsx(fieldsetLegendClasses))}>Calculation Details</legend> <div><label for="totalBirdhouses"${attr_class(clsx(labelBaseClasses))}>Total Number of Birdhouses to Place/Calculate:</label> <input type="number" id="totalBirdhouses"${attr("value", formState.totalBirdhouses)} min="1" step="1" max="999999" placeholder="e.g., 100"${attr_class(`${stringify(inputBaseClasses)} mt-1`)}/></div></fieldset> <button type="submit"${attr("disabled", isLoading, true)} class="w-full flex justify-center py-3 px-4 border border-transparent rounded-lg shadow-button text-sm font-semibold text-white bg-theme-accent hover:bg-theme-accent-hover focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-theme-card-bg focus:ring-theme-accent disabled:opacity-60 disabled:cursor-not-allowed transition-colors duration-150">${escape_html("Calculate Birdhouses")}</button></form>`;
  pop();
}
function _page($$payload, $$props) {
  push();
  const toolConfig = {
    name: "Birdhouse Run Calculator",
    description: "Estimate XP, nests, and valuable loot from your birdhouse runs based on log type and total houses.",
    iconSrc: "/images/birdhouse/redwood_bird_house.png"
  };
  $$payload.out += `<div class="max-w-6xl mx-auto"><div class="mb-6"><a href="/tools" class="inline-flex items-center text-sm font-medium text-theme-accent hover:text-theme-accent-hover transition-colors group"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-5 h-5 mr-1.5 transform transition-transform group-hover:-translate-x-1"><path fill-rule="evenodd" d="M17 10a.75.75 0 0 1-.75.75H5.56l2.72 2.72a.75.75 0 1 1-1.06 1.06l-4-4a.75.75 0 0 1 0-1.06l4-4a.75.75 0 0 1 1.06 1.06L5.56 9.25H16.25A.75.75 0 0 1 17 10Z" clip-rule="evenodd"></path></svg> Back to All Tools</a></div> <header class="mb-10 text-center"><div class="flex items-center justify-center mb-4"><div class="w-16 h-16 bg-gradient-to-r from-green-500 to-teal-600 rounded-xl flex items-center justify-center p-2 mr-4"><img${attr("src", toolConfig.iconSrc)}${attr("alt", `${stringify(toolConfig.name)} icon`)} class="w-full h-full object-contain"/></div> <h1 class="text-h1 text-theme-text-primary tracking-tight">${escape_html(toolConfig.name)}</h1></div> <p class="mt-3 text-lg text-theme-text-secondary">${escape_html(toolConfig.description)}</p></header> <div class="max-w-2xl mx-auto mb-8"><div class="bg-theme-bg-secondary rounded-card border border-theme-border-accent/20 p-6 shadow-card">`;
  InputForm($$payload);
  $$payload.out += `<!----></div></div> <div>`;
  {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--> `;
  {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--></div> <div class="max-w-4xl mx-auto mt-12"><div class="bg-theme-bg-secondary rounded-card border border-theme-border-accent/20 shadow-card overflow-hidden"><button class="w-full px-6 py-4 text-left bg-gradient-to-r from-green-500/10 via-teal-500/5 to-transparent hover:from-green-500/15 hover:via-teal-500/10 transition-all duration-200 flex items-center justify-between group"><div class="flex items-center space-x-3"><div class="w-8 h-8 bg-gradient-to-r from-green-500 to-teal-600 rounded-lg flex items-center justify-center flex-shrink-0"><svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg></div> <div><h3 class="text-lg font-semibold text-theme-text-primary">Calculation Methodology</h3> <p class="text-sm text-theme-text-secondary">Learn how Birdhouse XP rates and rewards are calculated</p></div></div> <svg${attr_class(`w-5 h-5 text-theme-text-secondary transition-transform duration-200 ${stringify("")}`)} fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg></button> `;
  {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--></div></div></div>`;
  pop();
}
export {
  _page as default
};
