import { c as create_ssr_component, v as validate_component, i as each, e as escape, b as spread, f as escape_object, d as escape_attribute_value } from './ssr-CwU1ZNF3.js';
import { M as Modal, L as Label, I as Input, B as Button } from './Modal-BD6gQ56w.js';
import { c as compute_rest_props } from './utils-Cb-NyS5y.js';
import { t as twMerge } from './bundle-mjs-BLGuyV0D.js';
import { T as Table, a as TableHead, b as TableHeadCell, c as TableBody, d as TableBodyRow, e as TableBodyCell } from './TableHeadCell-BPPFYRgE.js';
import { T as Tabs, a as TabItem } from './Tabs-hpKObKEs.js';
import { H as Heading } from './Heading-Bi-n-6bR.js';
import './client-BUusD8wq.js';
import './Wrapper-CKqRHixj.js';
import './Frame-CmaYSxgv.js';
import './ToolbarButton-D1qFMdlG.js';
import './index2-CpnJNRxb.js';
import './exports-BGi7-Rnc.js';

const Helper = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $$restProps = compute_rest_props($$props, ["helperClass", "color"]);
  let { helperClass = "text-xs font-normal text-gray-500 dark:text-gray-300" } = $$props;
  let { color = "gray" } = $$props;
  const colorClasses = {
    gray: "text-gray-900 dark:text-gray-300",
    green: "text-green-700 dark:text-green-500",
    red: "text-red-700 dark:text-red-500",
    disabled: "text-gray-400 dark:text-gray-500"
  };
  if ($$props.helperClass === void 0 && $$bindings.helperClass && helperClass !== void 0) $$bindings.helperClass(helperClass);
  if ($$props.color === void 0 && $$bindings.color && color !== void 0) $$bindings.color(color);
  return `<p${spread(
    [
      escape_object($$restProps),
      {
        class: escape_attribute_value(twMerge(helperClass, colorClasses[color], $$props.class))
      }
    ],
    {}
  )}>${slots.default ? slots.default({}) : ``}</p> `;
});
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { formModal = false } = $$props;
  let { form } = $$props;
  let { data } = $$props;
  if ($$props.formModal === void 0 && $$bindings.formModal && formModal !== void 0) $$bindings.formModal(formModal);
  if ($$props.form === void 0 && $$bindings.form && form !== void 0) $$bindings.form(form);
  if ($$props.data === void 0 && $$bindings.data && data !== void 0) $$bindings.data(data);
  let $$settled;
  let $$rendered;
  let previous_head = $$result.head;
  do {
    $$settled = true;
    $$result.head = previous_head;
    $$rendered = `${validate_component(Tabs, "Tabs").$$render(
      $$result,
      {
        class: "font-heading ml-2",
        contentClass: "",
        activeClasses: "p-4 text-theme border-b-2 border-theme"
      },
      {},
      {
        default: () => {
          return `${validate_component(TabItem, "TabItem").$$render($$result, { open: true, title: "Events" }, {}, {})} ${validate_component(TabItem, "TabItem").$$render($$result, { title: "Admin" }, {}, {})}`;
        }
      }
    )} <br> ${validate_component(Heading, "Heading").$$render(
      $$result,
      {
        tag: "h3",
        class: "font-heading ml-2",
        customSize: "text-xl"
      },
      {},
      {
        default: () => {
          return `Event Schedule`;
        }
      }
    )} <br> ${validate_component(Table, "Table").$$render(
      $$result,
      {
        hoverable: true,
        divClass: "ml-2 mr-2 !bg-theme font-default",
        color: "custom"
      },
      {},
      {
        default: () => {
          return `${validate_component(TableHead, "TableHead").$$render($$result, {}, {}, {
            default: () => {
              return `${validate_component(TableHeadCell, "TableHeadCell").$$render($$result, { class: "text-white" }, {}, {
                default: () => {
                  return `Time`;
                }
              })} ${validate_component(TableHeadCell, "TableHeadCell").$$render($$result, { class: "text-white" }, {}, {
                default: () => {
                  return `Division`;
                }
              })} ${validate_component(TableHeadCell, "TableHeadCell").$$render($$result, { class: "text-white" }, {}, {
                default: () => {
                  return `Status`;
                }
              })}`;
            }
          })} ${validate_component(TableBody, "TableBody").$$render($$result, { tableBodyClass: "divide-y" }, {}, {
            default: () => {
              return `${each(data.events, (event) => {
                return `${validate_component(TableBodyRow, "TableBodyRow").$$render($$result, { color: "default" }, {}, {
                  default: () => {
                    return `${validate_component(TableBodyCell, "TableBodyCell").$$render($$result, { class: "w-1 py-2" }, {}, {
                      default: () => {
                        return `<div class="text-black">${escape(event.time)}</div>`;
                      }
                    })} ${validate_component(TableBodyCell, "TableBodyCell").$$render($$result, { class: "w-1/3 py-2" }, {}, {
                      default: () => {
                        return `<div class="text-black">${escape(event.name)}</div>`;
                      }
                    })} ${validate_component(TableBodyCell, "TableBodyCell").$$render($$result, { class: "py-2" }, {}, {
                      default: () => {
                        return `${event.status == void 0 || event.status == 0 ? `<div class="text-red-500" data-svelte-h="svelte-1hl0tia">Not Started</div> ` : `${event.status == 1 || event.status == 2 ? `<div class="text-yellow-500" data-svelte-h="svelte-4tq761">In Progress</div> ` : `${event.status == 3 ? `<div class="text-green-500" data-svelte-h="svelte-13lvkj7">Complete</div> ` : ``}`}`}`;
                      }
                    })} `;
                  }
                })}`;
              })}`;
            }
          })}`;
        }
      }
    )} ${validate_component(Modal, "Modal").$$render(
      $$result,
      {
        size: "xs",
        autoclose: false,
        class: "w-full",
        open: formModal
      },
      {
        open: ($$value) => {
          formModal = $$value;
          $$settled = false;
        }
      },
      {
        default: () => {
          return `<form class="flex flex-col space-y-6" method="POST">${validate_component(Label, "Label").$$render($$result, { class: "space-y-2" }, {}, {
            default: () => {
              return `<span data-svelte-h="svelte-9difb2">Username</span> ${validate_component(Input, "Input").$$render(
                $$result,
                {
                  type: "text",
                  name: "username",
                  required: true,
                  class: "focus:ring-hover focus:border-hover"
                },
                {},
                {}
              )}`;
            }
          })} ${validate_component(Label, "Label").$$render($$result, { class: "space-y-2" }, {}, {
            default: () => {
              return `<span data-svelte-h="svelte-1kvjhoz">Password</span> ${validate_component(Input, "Input").$$render(
                $$result,
                {
                  class: "focus:ring-hover focus:border-hover",
                  type: "password",
                  name: "password",
                  required: true
                },
                {},
                {}
              )}`;
            }
          })} ${form?.incorrect ? `${validate_component(Helper, "Helper").$$render($$result, { class: "mt-2", color: "red" }, {}, {
            default: () => {
              return `Incorrect Credentials.`;
            }
          })}` : ``} ${validate_component(Button, "Button").$$render(
            $$result,
            {
              type: "submit",
              class: "w-full1 bg-theme hover:bg-hover"
            },
            {},
            {
              default: () => {
                return `Login`;
              }
            }
          )}</form>`;
        }
      }
    )}`;
  } while (!$$settled);
  return $$rendered;
});

export { Page as default };
//# sourceMappingURL=_page.svelte-DxFBY_Ac.js.map
