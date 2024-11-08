import { c as create_ssr_component, v as validate_component, i as each, e as escape } from './ssr-CwU1ZNF3.js';
import { T as Table, a as TableHead, b as TableHeadCell, c as TableBody, d as TableBodyRow, e as TableBodyCell } from './TableHeadCell-BPPFYRgE.js';
import { H as Heading } from './Heading-Bi-n-6bR.js';
import './utils-Cb-NyS5y.js';
import './bundle-mjs-BLGuyV0D.js';

const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { data } = $$props;
  if ($$props.data === void 0 && $$bindings.data && data !== void 0) $$bindings.data(data);
  return `${data.results == null ? `${validate_component(Heading, "Heading").$$render(
    $$result,
    {
      tag: "h5",
      class: "font-heading ml-2",
      customSize: "text-xl"
    },
    {},
    {
      default: () => {
        return `This division hasn&#39;t finished yet. Check back later!`;
      }
    }
  )}` : `${validate_component(Table, "Table").$$render($$result, { divClass: "ml-2 mr-2 font-default" }, {}, {
    default: () => {
      return `${validate_component(TableHead, "TableHead").$$render($$result, { class: "bg-theme text-white" }, {}, {
        default: () => {
          return `${validate_component(TableHeadCell, "TableHeadCell").$$render($$result, {}, {}, {
            default: () => {
              return `Rank`;
            }
          })} ${validate_component(TableHeadCell, "TableHeadCell").$$render($$result, {}, {}, {
            default: () => {
              return `Name`;
            }
          })}`;
        }
      })} ${validate_component(TableBody, "TableBody").$$render($$result, { tableBodyClass: "divide-y" }, {}, {
        default: () => {
          return `${each(data.results, (team) => {
            return `${validate_component(TableBodyRow, "TableBodyRow").$$render(
              $$result,
              {
                class: team.rank == 1 ? "bg-gold opacity-90" : team.rank == 2 ? "bg-silver opacity-90" : team.rank == 3 ? "bg-bronze opacity-90" : ""
              },
              {},
              {
                default: () => {
                  return `${validate_component(TableBodyCell, "TableBodyCell").$$render($$result, { class: "w-1 py-2" }, {}, {
                    default: () => {
                      return `<div class="${"text-lg " + escape(
                        team.rank === 2 || team.rank === 3 ? "text-white" : "text-black",
                        true
                      )}">${team.rank === 3 ? `${escape(team.rank)}T` : `${escape(team.rank)}`} </div>`;
                    }
                  })} ${validate_component(TableBodyCell, "TableBodyCell").$$render($$result, { class: "py-2" }, {}, {
                    default: () => {
                      return `<div class="${"text-lg " + escape(
                        team.rank === 2 || team.rank === 3 ? "text-white" : "text-black",
                        true
                      )}">${escape(team.name)} </div>`;
                    }
                  })} `;
                }
              }
            )}`;
          })}`;
        }
      })}`;
    }
  })}`}`;
});

export { Page as default };
//# sourceMappingURL=_page.svelte-Xeabzk8o.js.map
