import { Q as copy_payload, R as assign_payload, S as bind_props, C as pop, z as push, I as escape_html, G as attr_class, F as attr, T as clsx, E as ensure_array_like, J as stringify } from "../../../../chunks/index.js";
import { P as PlayerLookup } from "../../../../chunks/PlayerLookup.js";
const XP_THRESHOLDS = [
  0,
  83,
  174,
  276,
  388,
  512,
  650,
  801,
  969,
  1154,
  1358,
  1584,
  1833,
  2107,
  2411,
  2746,
  3115,
  3523,
  3973,
  4470,
  5018,
  5624,
  6291,
  7028,
  7842,
  8740,
  9730,
  10824,
  12031,
  13363,
  14833,
  16456,
  18247,
  20224,
  22406,
  24815,
  27473,
  30408,
  33648,
  37224,
  41171,
  45529,
  50339,
  55649,
  61512,
  67983,
  75127,
  83014,
  91721,
  101333,
  111945,
  123660,
  136594,
  150872,
  166636,
  184040,
  203254,
  224466,
  247886,
  273742,
  302288,
  333804,
  368599,
  407015,
  449428,
  496254,
  547953,
  605032,
  668051,
  737627,
  814445,
  899257,
  992895,
  1096278,
  1210421,
  1336443,
  1475581,
  1629200,
  1798808,
  1986068,
  2192818,
  2421087,
  2673114,
  2951373,
  3258594,
  3597792,
  3972294,
  4385776,
  4842295,
  5346332,
  5902831,
  6517253,
  7195629,
  7944614,
  8771558,
  9684577,
  10692629,
  11805606,
  13034431
];
const MAX_LEVEL = 99;
function getXpForLevel(level) {
  if (level < 1) return 0;
  if (level > MAX_LEVEL) level = MAX_LEVEL;
  if (level === 1) return 0;
  return XP_THRESHOLDS[level - 1];
}
function getLevelForXp(xp) {
  if (xp < 0) xp = 0;
  if (xp >= XP_THRESHOLDS[MAX_LEVEL - 1]) return MAX_LEVEL;
  for (let i = MAX_LEVEL - 1; i >= 0; i--) {
    if (xp >= XP_THRESHOLDS[i]) {
      return i + 1;
    }
  }
  return 1;
}
function _page($$payload, $$props) {
  push();
  let skillInfo, allTrainingMethods, currentActualLevel;
  let data = $$props["data"];
  let currentInputMode = "level";
  let currentLevelState = 1;
  let currentXPState = getXpForLevel(currentLevelState);
  let targetInputMode = "level";
  let targetLevelState = Math.min(currentLevelState + 1, 99);
  let targetXPState = getXpForLevel(targetLevelState);
  let playerUsername = "";
  let playerLookupLoading = false;
  let calculationError = null;
  function syncCurrentXP() {
    currentXPState = getXpForLevel(currentLevelState);
  }
  function syncTargetXP() {
    targetXPState = getXpForLevel(targetLevelState);
  }
  function calculateMethodMetrics(method, currentXP, targetXP) {
    if (targetXP <= currentXP || !skillInfo) return null;
    const totalXPToGain = targetXP - currentXP;
    let timeToTargetHours = void 0;
    let actionsNeeded = void 0;
    let marksOfGraceEarned = void 0;
    if (method.xpRate > 0) {
      timeToTargetHours = totalXPToGain / method.xpRate;
    }
    if (method.xpPerAction && method.xpPerAction > 0) {
      actionsNeeded = Math.ceil(totalXPToGain / method.xpPerAction);
    }
    if (skillInfo.skillNameCanonical === "agility") {
      if (method.marksPerHour !== void 0 && timeToTargetHours !== void 0) {
        marksOfGraceEarned = method.marksPerHour * timeToTargetHours;
      }
    }
    return {
      actionsNeeded,
      timeToTargetHours,
      timeToTargetFormatted: formatHours(timeToTargetHours),
      marksOfGraceEarned
    };
  }
  function formatHours(hours) {
    if (hours === void 0 || isNaN(hours) || !isFinite(hours)) return void 0;
    if (hours === 0 && !(hours > 0)) return "0h 0m";
    const totalMinutes = Math.round(hours * 60);
    if (totalMinutes === 0 && hours > 0) return "<1m";
    const h = Math.floor(totalMinutes / 60);
    const m = totalMinutes % 60;
    return `${h}h ${m}m`;
  }
  function formatNum(num, decimals = 0) {
    if (num === null || num === void 0 || typeof num !== "number" || isNaN(num)) return "N/A";
    return num.toLocaleString(void 0, {
      minimumFractionDigits: decimals,
      maximumFractionDigits: decimals
    });
  }
  const inputBaseClasses = "block w-full rounded-md border-0 py-2 px-3.5 bg-gray-700/50 text-theme-text-primary shadow-sm ring-1 ring-inset ring-theme-border-input placeholder:text-theme-text-tertiary focus:ring-2 focus:ring-inset focus:ring-theme-accent sm:text-sm sm:leading-6 shadow-inner-border transition-colors duration-150";
  const radioLabelClasses = "flex items-center text-sm text-theme-text-secondary cursor-pointer";
  const radioInputClasses = "h-4 w-4 border-theme-border-input text-theme-accent focus:ring-theme-accent focus:ring-offset-theme-card-bg";
  skillInfo = data.skillData;
  allTrainingMethods = data.skillData?.trainingMethods.sort((a, b) => a.levelReq - b.levelReq || b.xpRate - a.xpRate) || [];
  syncCurrentXP();
  syncTargetXP();
  {
    if (targetXPState <= currentXPState) {
      calculationError = "Target must be greater than current progress.";
    } else {
      calculationError = null;
    }
  }
  currentActualLevel = getLevelForXp(currentXPState);
  let $$settled = true;
  let $$inner_payload;
  function $$render_inner($$payload2) {
    $$payload2.out += `<div class="max-w-6xl mx-auto space-y-10 py-8 px-4 sm:px-6 lg:px-8"><div class="mb-6"><a href="/skills" class="inline-flex items-center text-sm font-medium text-theme-accent hover:text-theme-accent-hover transition-colors group"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-5 h-5 mr-1.5 transform transition-transform group-hover:-translate-x-1"><path fill-rule="evenodd" d="M17 10a.75.75 0 0 1-.75.75H5.56l2.72 2.72a.75.75 0 1 1-1.06 1.06l-4-4a.75.75 0 0 1 0-1.06l4-4a.75.75 0 0 1 1.06 1.06L5.56 9.25H16.25A.75.75 0 0 1 17 10Z" clip-rule="evenodd"></path></svg> Back to All Skills</a></div> <header class="text-center">`;
    if (skillInfo) {
      $$payload2.out += "<!--[-->";
      $$payload2.out += `<h1 class="text-h1 text-theme-text-primary tracking-tight">${escape_html(skillInfo.skillNameDisplay)} Calculator</h1> `;
      if (skillInfo.description) {
        $$payload2.out += "<!--[-->";
        $$payload2.out += `<p class="mt-3 text-lg text-theme-text-secondary max-w-2xl mx-auto">${escape_html(skillInfo.description)}</p>`;
      } else {
        $$payload2.out += "<!--[!-->";
      }
      $$payload2.out += `<!--]-->`;
    } else {
      $$payload2.out += "<!--[!-->";
      $$payload2.out += `<h1 class="text-h1 text-theme-text-primary tracking-tight">Skill Calculator</h1> <p class="mt-3 text-lg text-theme-text-secondary max-w-2xl mx-auto">Loading skill data...</p>`;
    }
    $$payload2.out += `<!--]--></header> `;
    if (skillInfo) {
      $$payload2.out += "<!--[-->";
      $$payload2.out += `<section class="bg-theme-card-bg shadow-card rounded-lg border border-theme-border p-6 space-y-8"><h2 class="text-h3 text-theme-text-primary border-b border-theme-border pb-3">Your Progress &amp; Goals</h2> <div class="space-y-4"><div class="block text-sm font-semibold text-theme-text-primary mb-3">Player Lookup (Optional)</div> <div>`;
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
      $$payload2.out += `<!----> <p class="text-theme-text-tertiary text-xs mt-1">Automatically loads your current ${escape_html(skillInfo?.skillNameDisplay || "skill")} level from OSRS hiscores</p></div></div> <div class="grid grid-cols-1 md:grid-cols-2 gap-x-8 gap-y-6"><div class="space-y-3"><p class="text-base font-semibold text-theme-text-primary">Current ${escape_html(skillInfo.skillNameDisplay)}</p> <div class="flex items-center space-x-4"><label${attr_class(clsx(radioLabelClasses))}><input type="radio"${attr("checked", currentInputMode === "level", true)} value="level" name="currentInputMode"${attr_class(clsx(radioInputClasses))}/> <span class="ml-2">Level</span></label> <label${attr_class(clsx(radioLabelClasses))}><input type="radio"${attr("checked", currentInputMode === "xp", true)} value="xp" name="currentInputMode"${attr_class(clsx(radioInputClasses))}/> <span class="ml-2">XP</span></label></div> `;
      {
        $$payload2.out += "<!--[-->";
        $$payload2.out += `<input type="number" placeholder="Current Level"${attr("value", currentLevelState)} min="1" max="99" step="1"${attr_class(clsx(inputBaseClasses))}/> <p class="text-xs text-theme-text-tertiary pl-1">XP: ${escape_html(formatNum(currentXPState))}</p>`;
      }
      $$payload2.out += `<!--]--></div> <div class="space-y-3"><p class="text-base font-semibold text-theme-text-primary">Target ${escape_html(skillInfo.skillNameDisplay)}</p> <div class="flex items-center space-x-4"><label${attr_class(clsx(radioLabelClasses))}><input type="radio"${attr("checked", targetInputMode === "level", true)} value="level" name="targetInputMode"${attr_class(clsx(radioInputClasses))}/> <span class="ml-2">Level</span></label> <label${attr_class(clsx(radioLabelClasses))}><input type="radio"${attr("checked", targetInputMode === "xp", true)} value="xp" name="targetInputMode"${attr_class(clsx(radioInputClasses))}/> <span class="ml-2">XP</span></label></div> `;
      {
        $$payload2.out += "<!--[-->";
        $$payload2.out += `<input type="number" placeholder="Target Level"${attr("value", targetLevelState)}${attr("min", currentLevelState + 1)} max="99" step="1"${attr_class(clsx(inputBaseClasses))}/> <p class="text-xs text-theme-text-tertiary pl-1">XP: ${escape_html(formatNum(targetXPState))}</p>`;
      }
      $$payload2.out += `<!--]--></div></div></section> `;
      if (calculationError) {
        $$payload2.out += "<!--[-->";
        $$payload2.out += `<div class="bg-red-900/80 border border-red-700 text-red-100 px-4 py-3 rounded-lg shadow-md text-sm" role="alert"><div class="flex items-center"><svg class="fill-current h-5 w-5 text-red-400 mr-3" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M2.93 17.07A10 10 0 1 1 17.07 2.93 10 10 0 0 1 2.93 17.07zM11.414 10l2.829-2.829a1 1 0 0 0-1.414-1.414L10 8.586 7.172 5.757a1 1 0 0 0-1.414 1.414L8.586 10l-2.829 2.829a1 1 0 1 0 1.414 1.414L10 11.414l2.829 2.829a1 1 0 0 0 1.414-1.414L11.414 10z"></path></svg> <p>${escape_html(calculationError)}</p></div></div>`;
      } else {
        $$payload2.out += "<!--[!-->";
      }
      $$payload2.out += `<!--]--> <section class="bg-theme-card-bg shadow-card rounded-lg border border-theme-border p-6 space-y-8 overflow-hidden"><h2 class="text-h3 text-theme-text-primary border-b border-theme-border pb-3 mb-0 px-6 sm:px-8 pt-6 sm:pt-8">Training Methods for ${escape_html(skillInfo.skillNameDisplay)}</h2> `;
      if (allTrainingMethods.length > 0) {
        $$payload2.out += "<!--[-->";
        const each_array = ensure_array_like(allTrainingMethods);
        $$payload2.out += `<div class="overflow-x-auto m-10"><table class="min-w-full divide-y divide-theme-border-subtle text-sm"><thead class="bg-gray-700/30 sticky top-0 z-10"><tr><th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary">Method</th><th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary">Lvl</th><th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary">XP/hr</th>`;
        if (skillInfo.skillNameCanonical === "agility") {
          $$payload2.out += "<!--[-->";
          $$payload2.out += `<th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary">Marks/hr</th>`;
        } else {
          $$payload2.out += "<!--[!-->";
        }
        $$payload2.out += `<!--]--><th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary">Actions to Target</th><th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary">Time to Target</th>`;
        if (skillInfo.skillNameCanonical === "agility") {
          $$payload2.out += "<!--[-->";
          $$payload2.out += `<th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary">Total Marks to Target</th>`;
        } else {
          $$payload2.out += "<!--[!-->";
        }
        $$payload2.out += `<!--]--><th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary hidden md:table-cell">Type</th><th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary hidden lg:table-cell">Location</th></tr></thead><tbody class="divide-y divide-theme-border-subtle/50"><!--[-->`;
        for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
          let method = each_array[$$index];
          const isAvailable = currentActualLevel >= method.levelReq;
          const methodCalcs = calculateMethodMetrics(method, currentXPState, targetXPState);
          const colspanForNotes = 3 + (skillInfo.skillNameCanonical === "agility" ? 1 : 0) + 2 + (skillInfo.skillNameCanonical === "agility" ? 1 : 0) + 1 + 1;
          $$payload2.out += `<tr${attr_class("hover:bg-gray-700/10", void 0, {
            "opacity-60": !isAvailable,
            "pointer-events-none": !isAvailable
          })}><td${attr_class(`whitespace-nowrap px-3 py-3 font-medium ${stringify(isAvailable ? "text-theme-text-primary" : "text-theme-text-tertiary")}`)}>${escape_html(method.name)}</td><td${attr_class(`whitespace-nowrap px-3 py-3 ${stringify(isAvailable ? "text-theme-text-secondary" : "text-theme-text-tertiary")}`)}>${escape_html(method.levelReq)}</td><td${attr_class(`whitespace-nowrap px-3 py-3 ${stringify(isAvailable ? "text-theme-text-secondary" : "text-theme-text-tertiary")}`)}>${escape_html(formatNum(method.xpRate))}</td>`;
          if (skillInfo.skillNameCanonical === "agility") {
            $$payload2.out += "<!--[-->";
            $$payload2.out += `<td${attr_class(`whitespace-nowrap px-3 py-3 ${stringify(isAvailable ? "text-theme-text-secondary" : "text-theme-text-tertiary")}`)}>${escape_html(method.marksPerHour ? formatNum(method.marksPerHour, 1) : "N/A")}</td>`;
          } else {
            $$payload2.out += "<!--[!-->";
          }
          $$payload2.out += `<!--]--><td${attr_class(`whitespace-nowrap px-3 py-3 ${stringify(isAvailable ? "text-theme-text-secondary" : "text-theme-text-tertiary")}`)}>${escape_html(methodCalcs && methodCalcs.actionsNeeded !== void 0 ? formatNum(methodCalcs.actionsNeeded) : targetXPState <= currentXPState ? "-" : "N/A")}</td><td${attr_class(`whitespace-nowrap px-3 py-3 ${stringify(isAvailable ? "text-theme-text-secondary" : "text-theme-text-tertiary")}`)}>${escape_html(methodCalcs?.timeToTargetFormatted || (targetXPState <= currentXPState ? "-" : "N/A"))}</td>`;
          if (skillInfo.skillNameCanonical === "agility") {
            $$payload2.out += "<!--[-->";
            $$payload2.out += `<td${attr_class(`whitespace-nowrap px-3 py-3 ${stringify(isAvailable ? "text-theme-text-secondary" : "text-theme-text-tertiary")}`)}>${escape_html(methodCalcs && methodCalcs.marksOfGraceEarned !== void 0 ? formatNum(methodCalcs.marksOfGraceEarned, 1) : targetXPState <= currentXPState ? "-" : "N/A")}</td>`;
          } else {
            $$payload2.out += "<!--[!-->";
          }
          $$payload2.out += `<!--]--><td${attr_class(`whitespace-nowrap px-3 py-3 ${stringify(isAvailable ? "text-theme-text-secondary" : "text-theme-text-tertiary")} hidden md:table-cell`)}>${escape_html(method.type || "N/A")}</td><td${attr_class(`px-3 py-3 ${stringify(isAvailable ? "text-theme-text-secondary" : "text-theme-text-tertiary")} hidden lg:table-cell min-w-[150px]`)}>${escape_html(method.location || "N/A")}</td></tr> `;
          if (method.notes || method.itemsRequired && method.itemsRequired.length > 0 || method.questsRequired && method.questsRequired.length > 0) {
            $$payload2.out += "<!--[-->";
            $$payload2.out += `<tr${attr_class("bg-gray-700/10 hover:bg-gray-700/20", void 0, { "opacity-60": !isAvailable })}><td${attr_class(`px-3 py-2 italic text-xs ${stringify(isAvailable ? "text-theme-text-tertiary" : "text-gray-500")}`)}${attr("colspan", colspanForNotes)}>`;
            if (method.itemsRequired && method.itemsRequired.length > 0) {
              $$payload2.out += "<!--[-->";
              $$payload2.out += `<strong>Items:</strong> ${escape_html(method.itemsRequired.join(", "))}`;
            } else {
              $$payload2.out += "<!--[!-->";
            }
            $$payload2.out += `<!--]--> `;
            if (method.questsRequired && method.questsRequired.length > 0) {
              $$payload2.out += "<!--[-->";
              $$payload2.out += `<br${attr_class("", void 0, {
                "mt-1": method.itemsRequired && method.itemsRequired.length > 0
              })}/><strong>Quests:</strong> ${escape_html(method.questsRequired.join(", "))}`;
            } else {
              $$payload2.out += "<!--[!-->";
            }
            $$payload2.out += `<!--]--> `;
            if (method.notes) {
              $$payload2.out += "<!--[-->";
              $$payload2.out += `<br${attr_class("", void 0, {
                "mt-1": method.itemsRequired && method.itemsRequired.length > 0 || method.questsRequired && method.questsRequired.length > 0
              })}/> <em>${escape_html(method.notes)}</em>`;
            } else {
              $$payload2.out += "<!--[!-->";
            }
            $$payload2.out += `<!--]--></td></tr>`;
          } else {
            $$payload2.out += "<!--[!-->";
          }
          $$payload2.out += `<!--]-->`;
        }
        $$payload2.out += `<!--]--></tbody></table></div>`;
      } else {
        $$payload2.out += "<!--[!-->";
        $$payload2.out += `<p class="text-theme-text-secondary px-6 sm:px-8">No training methods available for ${escape_html(skillInfo.skillNameDisplay)}.</p>`;
      }
      $$payload2.out += `<!--]--></section>`;
    } else {
      $$payload2.out += "<!--[!-->";
      $$payload2.out += `<div class="text-center py-10"><p class="text-theme-text-secondary text-lg">Skill data not loaded or skill not found. Check backend and API path.</p></div>`;
    }
    $$payload2.out += `<!--]--></div>`;
  }
  do {
    $$settled = true;
    $$inner_payload = copy_payload($$payload);
    $$render_inner($$inner_payload);
  } while (!$$settled);
  assign_payload($$payload, $$inner_payload);
  bind_props($$props, { data });
  pop();
}
export {
  _page as default
};
