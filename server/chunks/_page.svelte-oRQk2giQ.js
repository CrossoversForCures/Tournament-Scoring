import { s as subscribe } from './utils-Cb-NyS5y.js';
import { c as create_ssr_component, v as validate_component, i as each, e as escape } from './ssr-CwU1ZNF3.js';
import { T as Table, a as TableHead, b as TableHeadCell, c as TableBody, d as TableBodyRow, e as TableBodyCell } from './TableHeadCell-BPPFYRgE.js';
import { H as Heading } from './Heading-Bi-n-6bR.js';
import { i as isAdmin } from './admin-DrRi-8BF.js';
import './client-BUusD8wq.js';
import './bundle-mjs-BLGuyV0D.js';
import './index2-CpnJNRxb.js';
import './exports-BGi7-Rnc.js';

const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $isAdmin, $$unsubscribe_isAdmin;
  $$unsubscribe_isAdmin = subscribe(isAdmin, (value) => $isAdmin = value);
  let { data } = $$props;
  if ($$props.data === void 0 && $$bindings.data && data !== void 0) $$bindings.data(data);
  $$unsubscribe_isAdmin();
  return `${data.seeding == null ? `${$isAdmin ? `<form method="POST" action="?/start">${validate_component(Heading, "Heading").$$render(
    $$result,
    {
      tag: "h5",
      class: "font-heading ml-2",
      customSize: "text-xl"
    },
    {},
    {
      default: () => {
        return `The elimination round for this division hasn&#39;t started yet.
				<button class="link text-theme hover:text-hover" type="submit" data-svelte-h="svelte-f9gi9z">Start Elimination?</button>`;
      }
    }
  )}</form>` : `${validate_component(Heading, "Heading").$$render(
    $$result,
    {
      tag: "h5",
      class: "font-heading ml-2",
      customSize: "text-xl"
    },
    {},
    {
      default: () => {
        return `The elimination round for this division hasn&#39;t started yet. Check back later!`;
      }
    }
  )}`}` : `${validate_component(Table, "Table").$$render($$result, { divClass: "ml-2 mr-2 font-default" }, {}, {
    default: () => {
      return `${validate_component(TableHead, "TableHead").$$render($$result, { class: "bg-theme text-white" }, {}, {
        default: () => {
          return `${validate_component(TableHeadCell, "TableHeadCell").$$render($$result, {}, {}, {
            default: () => {
              return `Seed`;
            }
          })} ${validate_component(TableHeadCell, "TableHeadCell").$$render($$result, {}, {}, {
            default: () => {
              return `Team`;
            }
          })} ${validate_component(TableHeadCell, "TableHeadCell").$$render($$result, {}, {}, {
            default: () => {
              return `Games Won`;
            }
          })} ${validate_component(TableHeadCell, "TableHeadCell").$$render($$result, {}, {}, {
            default: () => {
              return `Total Points`;
            }
          })}`;
        }
      })} ${validate_component(TableBody, "TableBody").$$render($$result, {}, {}, {
        default: () => {
          return `${each(data.seeding, (team) => {
            return `${validate_component(TableBodyRow, "TableBodyRow").$$render($$result, { color: "custom" }, {}, {
              default: () => {
                return `${validate_component(TableBodyCell, "TableBodyCell").$$render($$result, { class: "w-1 py-2" }, {}, {
                  default: () => {
                    return `<div class="text-black">${escape(team.seeding)}</div>`;
                  }
                })} ${validate_component(TableBodyCell, "TableBodyCell").$$render($$result, { class: "w-1/3 py-2" }, {}, {
                  default: () => {
                    return `<div class="text-black">${escape(team.name)}</div>`;
                  }
                })} ${validate_component(TableBodyCell, "TableBodyCell").$$render($$result, { class: "w-1/4 py-2" }, {}, {
                  default: () => {
                    return `<div class="text-black">${team.poolsWon === void 0 ? `0` : `${escape(team.poolsWon)}`} </div>`;
                  }
                })} ${validate_component(TableBodyCell, "TableBodyCell").$$render($$result, { class: "w-1/4 py-2" }, {}, {
                  default: () => {
                    return `<div class="text-black">${team.totalPoints === void 0 ? `0` : `${escape(team.totalPoints)}`} </div>`;
                  }
                })} `;
              }
            })}`;
          })}`;
        }
      })}`;
    }
  })}`}`;
});

export { Page as default };
//# sourceMappingURL=_page.svelte-oRQk2giQ.js.map
