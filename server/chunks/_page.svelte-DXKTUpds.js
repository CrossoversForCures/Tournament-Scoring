import { s as subscribe } from './utils-Cb-NyS5y.js';
import { c as create_ssr_component, v as validate_component, i as each, e as escape, a as add_attribute } from './ssr-CwU1ZNF3.js';
import { w as writable } from './index2-CpnJNRxb.js';
import { M as Modal, L as Label, I as Input, B as Button } from './Modal-BD6gQ56w.js';
import { T as Table, a as TableHead, b as TableHeadCell, c as TableBody, d as TableBodyRow, e as TableBodyCell } from './TableHeadCell-BPPFYRgE.js';
import { H as Heading } from './Heading-Bi-n-6bR.js';
import { E as EditOutline } from './EditOutline-B6i7YDYF.js';
import './client-BUusD8wq.js';
import { i as isAdmin } from './admin-DrRi-8BF.js';
import './bundle-mjs-BLGuyV0D.js';
import './Wrapper-CKqRHixj.js';
import './Frame-CmaYSxgv.js';
import './ToolbarButton-D1qFMdlG.js';
import './exports-BGi7-Rnc.js';

const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $isAdmin, $$unsubscribe_isAdmin;
  let $editingGame, $$unsubscribe_editingGame;
  let $editingTeam1, $$unsubscribe_editingTeam1;
  let $editingTeam2, $$unsubscribe_editingTeam2;
  $$unsubscribe_isAdmin = subscribe(isAdmin, (value) => $isAdmin = value);
  let { data } = $$props;
  let { formModal = false } = $$props;
  const editingGame = writable(null);
  $$unsubscribe_editingGame = subscribe(editingGame, (value) => $editingGame = value);
  const editingTeam1 = writable(null);
  $$unsubscribe_editingTeam1 = subscribe(editingTeam1, (value) => $editingTeam1 = value);
  const editingTeam2 = writable(null);
  $$unsubscribe_editingTeam2 = subscribe(editingTeam2, (value) => $editingTeam2 = value);
  if ($$props.data === void 0 && $$bindings.data && data !== void 0) $$bindings.data(data);
  if ($$props.formModal === void 0 && $$bindings.formModal && formModal !== void 0) $$bindings.formModal(formModal);
  let $$settled;
  let $$rendered;
  let previous_head = $$result.head;
  do {
    $$settled = true;
    $$result.head = previous_head;
    $$rendered = `${data.games == null ? `${$isAdmin ? `<form method="POST" action="?/start">${validate_component(Heading, "Heading").$$render(
      $$result,
      {
        tag: "h5",
        class: "font-heading ml-2",
        customSize: "text-xl"
      },
      {},
      {
        default: () => {
          return `This division hasn&#39;t started yet.
				<button class="link text-theme hover:text-hover" type="submit" data-svelte-h="svelte-ciiku1">Start Pools?</button>`;
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
          return `This division hasn&#39;t started yet. Check back later!`;
        }
      }
    )}`}` : `${each(Object.keys(data.games), (round) => {
      return `${validate_component(Heading, "Heading").$$render(
        $$result,
        {
          tag: "h5",
          class: "font-heading ml-2",
          customSize: "text-xl"
        },
        {},
        {
          default: () => {
            return `Round ${escape(round)}`;
          }
        }
      )} ${validate_component(Table, "Table").$$render($$result, { divClass: "ml-2 mr-2 font-default" }, {}, {
        default: () => {
          return `${validate_component(TableHead, "TableHead").$$render($$result, { class: "bg-theme text-white" }, {}, {
            default: () => {
              return `${validate_component(TableHeadCell, "TableHeadCell").$$render($$result, {}, {}, {
                default: () => {
                  return `Court`;
                }
              })} ${validate_component(TableHeadCell, "TableHeadCell").$$render($$result, {}, {}, {
                default: () => {
                  return `Team 1`;
                }
              })} ${validate_component(TableHeadCell, "TableHeadCell").$$render($$result, {}, {}, {
                default: () => {
                  return `Score`;
                }
              })} ${validate_component(TableHeadCell, "TableHeadCell").$$render($$result, {}, {}, {
                default: () => {
                  return `Team 2`;
                }
              })} ${validate_component(TableHeadCell, "TableHeadCell").$$render($$result, {}, {}, {
                default: () => {
                  return `Score`;
                }
              })} ${$isAdmin ? `${validate_component(TableHeadCell, "TableHeadCell").$$render($$result, {}, {}, {
                default: () => {
                  return `Update`;
                }
              })}` : ``} `;
            }
          })} ${validate_component(TableBody, "TableBody").$$render($$result, {}, {}, {
            default: () => {
              return `${each(data.games[round], (game) => {
                return `${validate_component(TableBodyRow, "TableBodyRow").$$render($$result, { color: "default" }, {}, {
                  default: () => {
                    return `${validate_component(TableBodyCell, "TableBodyCell").$$render($$result, { class: "w-1/6 py-2" }, {}, {
                      default: () => {
                        return `Court ${escape(game.court)}`;
                      }
                    })} ${validate_component(TableBodyCell, "TableBodyCell").$$render($$result, { class: "w-1/4 py-2" }, {}, {
                      default: () => {
                        return `${escape(game.team1Name)}`;
                      }
                    })} ${validate_component(TableBodyCell, "TableBodyCell").$$render(
                      $$result,
                      {
                        class: "w-1/6 py-2 font-semibold " + (game.team1Score > game.team2Score ? "text-green-500" : "text-red-500")
                      },
                      {},
                      {
                        default: () => {
                          return `${escape(game.team1Score < 0 ? "" : game.team1Score)} `;
                        }
                      }
                    )} ${validate_component(TableBodyCell, "TableBodyCell").$$render($$result, { class: "w-1/4 py-2" }, {}, {
                      default: () => {
                        return `${escape(game.team2Name)}`;
                      }
                    })} ${validate_component(TableBodyCell, "TableBodyCell").$$render(
                      $$result,
                      {
                        class: "w-1/6 py-2 font-semibold " + (game.team2Score > game.team1Score ? "text-green-500" : "text-red-500")
                      },
                      {},
                      {
                        default: () => {
                          return `${escape(game.team2Score < 0 ? "" : game.team2Score)} `;
                        }
                      }
                    )} ${$isAdmin ? `${validate_component(TableBodyCell, "TableBodyCell").$$render($$result, { class: "px-6 py-0" }, {}, {
                      default: () => {
                        return `<button>${validate_component(EditOutline, "EditOutline").$$render(
                          $$result,
                          {
                            class: "text-theme h-7 w-7 content-center "
                          },
                          {},
                          {}
                        )}</button> `;
                      }
                    })}` : ``} ${validate_component(Modal, "Modal").$$render(
                      $$result,
                      {
                        size: "xs",
                        autoclose: false,
                        class: "w-full",
                        backdropClass: "fixed inset-0 z-40 bg-gray-900 !opacity-10",
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
                          return `<form class="flex flex-col space-y-6" method="POST" action="?/update"><input type="hidden" name="gameId"${add_attribute("value", $editingGame, 0)}> ${validate_component(Label, "Label").$$render($$result, { class: "active:border-theme space-y-2" }, {}, {
                            default: () => {
                              return `<span>Team 1 (${escape($editingTeam1)}) Score</span> ${validate_component(Input, "Input").$$render(
                                $$result,
                                {
                                  type: "number",
                                  name: "team1Score",
                                  required: true
                                },
                                {},
                                {}
                              )} `;
                            }
                          })} ${validate_component(Label, "Label").$$render($$result, { class: "space-y-2" }, {}, {
                            default: () => {
                              return `<span>Team 2 (${escape($editingTeam2)}) Score</span> ${validate_component(Input, "Input").$$render(
                                $$result,
                                {
                                  type: "number",
                                  name: "team2Score",
                                  required: true
                                },
                                {},
                                {}
                              )} `;
                            }
                          })} ${validate_component(Button, "Button").$$render(
                            $$result,
                            {
                              type: "submit",
                              class: "w-full1 bg-theme hover:bg-hover"
                            },
                            {},
                            {
                              default: () => {
                                return `Confirm`;
                              }
                            }
                          )}</form> `;
                        }
                      }
                    )} `;
                  }
                })}`;
              })} `;
            }
          })} `;
        }
      })} <br>`;
    })}`}`;
  } while (!$$settled);
  $$unsubscribe_isAdmin();
  $$unsubscribe_editingGame();
  $$unsubscribe_editingTeam1();
  $$unsubscribe_editingTeam2();
  return $$rendered;
});

export { Page as default };
//# sourceMappingURL=_page.svelte-DXKTUpds.js.map
