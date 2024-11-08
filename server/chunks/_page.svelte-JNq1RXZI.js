import { c as create_ssr_component, v as validate_component, i as each, e as escape } from './ssr-CwU1ZNF3.js';
import { T as Table, a as TableHead, b as TableHeadCell, c as TableBody, d as TableBodyRow, e as TableBodyCell } from './TableHeadCell-BPPFYRgE.js';
import { H as Heading } from './Heading-Bi-n-6bR.js';
import './utils-Cb-NyS5y.js';
import './bundle-mjs-BLGuyV0D.js';

const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { data } = $$props;
  if ($$props.data === void 0 && $$bindings.data && data !== void 0) $$bindings.data(data);
  return `${data.teams === null ? `${validate_component(Heading, "Heading").$$render(
    $$result,
    {
      tag: "h5",
      class: "font-heading ml-2",
      customSize: "text-xl"
    },
    {},
    {
      default: () => {
        return `There are no teams registered for this division yet. Check back later!`;
      }
    }
  )}` : `${validate_component(Table, "Table").$$render($$result, { divClass: "ml-2 mr-2 font-default" }, {}, {
    default: () => {
      return `${validate_component(TableHead, "TableHead").$$render($$result, { class: "bg-theme text-white" }, {}, {
        default: () => {
          return `${validate_component(TableHeadCell, "TableHeadCell").$$render($$result, {}, {}, {
            default: () => {
              return `Name`;
            }
          })}`;
        }
      })} ${validate_component(TableBody, "TableBody").$$render($$result, { tableBodyClass: "divide-y" }, {}, {
        default: () => {
          return `${each(data.teams, (team) => {
            return `${validate_component(TableBodyRow, "TableBodyRow").$$render($$result, { color: "default" }, {}, {
              default: () => {
                return `${validate_component(TableBodyCell, "TableBodyCell").$$render($$result, { class: "py-2" }, {}, {
                  default: () => {
                    return `<div class="text-black">${escape(team.name)}</div>`;
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
//# sourceMappingURL=_page.svelte-JNq1RXZI.js.map
