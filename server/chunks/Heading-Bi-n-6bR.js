import { c as compute_rest_props } from './utils-Cb-NyS5y.js';
import { c as create_ssr_component, b as spread, f as escape_object, d as escape_attribute_value } from './ssr-CwU1ZNF3.js';
import { t as twMerge } from './bundle-mjs-BLGuyV0D.js';

const void_element_names = /^(?:area|base|br|col|command|embed|hr|img|input|keygen|link|meta|param|source|track|wbr)$/;
function is_void(name) {
  return void_element_names.test(name) || name.toLowerCase() === "!doctype";
}
const Heading = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $$restProps = compute_rest_props($$props, ["tag", "color", "customSize"]);
  let { tag = "h1" } = $$props;
  let { color = "text-gray-900 dark:text-white" } = $$props;
  let { customSize = "" } = $$props;
  const textSizes = {
    h1: "text-5xl font-extrabold",
    h2: "text-4xl font-bold",
    h3: "text-3xl font-bold",
    h4: "text-2xl font-bold",
    h5: "text-xl font-bold",
    h6: "text-lg font-bold"
  };
  if ($$props.tag === void 0 && $$bindings.tag && tag !== void 0) $$bindings.tag(tag);
  if ($$props.color === void 0 && $$bindings.color && color !== void 0) $$bindings.color(color);
  if ($$props.customSize === void 0 && $$bindings.customSize && customSize !== void 0) $$bindings.customSize(customSize);
  return `${((tag$1) => {
    return tag$1 ? `<${tag}${spread(
      [
        escape_object($$restProps),
        {
          class: escape_attribute_value(twMerge(customSize ? customSize : textSizes[tag], color, "w-full", $$props.class))
        }
      ],
      {}
    )}>${is_void(tag$1) ? "" : `${slots.default ? slots.default({}) : ``}`}${is_void(tag$1) ? "" : `</${tag$1}>`}` : "";
  })(tag)} `;
});

export { Heading as H, is_void as i };
//# sourceMappingURL=Heading-Bi-n-6bR.js.map
