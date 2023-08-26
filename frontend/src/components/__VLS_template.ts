import DefaultLayout from "./DefaultLayout.vue";
import { __VLS_publicComponent, __VLS_internalComponent, __VLS_componentsOption, __VLS_name } from "./Index.vue";

function __VLS_template() {
let __VLS_ctx!: InstanceType<__VLS_PickNotAny<typeof __VLS_publicComponent, new () => {}>> & InstanceType<__VLS_PickNotAny<typeof __VLS_internalComponent, new () => {}>> & {};
/* Components */
let __VLS_otherComponents!: NonNullable<typeof __VLS_internalComponent extends { components: infer C; } ? C : {}> & typeof __VLS_componentsOption;
let __VLS_own!: __VLS_SelfComponent<typeof __VLS_name, typeof __VLS_internalComponent & typeof __VLS_publicComponent & (new () => { $slots: typeof __VLS_slots; })>;
let __VLS_localComponents!: typeof __VLS_otherComponents & Omit<typeof __VLS_own, keyof typeof __VLS_otherComponents>;
let __VLS_components!: typeof __VLS_localComponents & __VLS_GlobalComponents & typeof __VLS_ctx;
/* Style Scoped */
type __VLS_StyleScopedClasses = {};
let __VLS_styleScopedClasses!: __VLS_StyleScopedClasses | keyof __VLS_StyleScopedClasses | (keyof __VLS_StyleScopedClasses)[];
/* CSS variable injection */
/* CSS variable injection end */
let __VLS_resolvedLocalAndGlobalComponents!: {} &
__VLS_WithComponent<'DefaultLayout', typeof __VLS_localComponents, "DefaultLayout", "DefaultLayout", "DefaultLayout"> &
__VLS_WithComponent<'Post', typeof __VLS_localComponents, "Post", "Post", "Post">;
__VLS_components.DefaultLayout; __VLS_components.DefaultLayout;
// @ts-ignore
[DefaultLayout, DefaultLayout,];
({} as __VLS_IntrinsicElements).p; ({} as __VLS_IntrinsicElements).p;
__VLS_components.Post;
// @ts-ignore
[Post,];
{
let __VLS_0!: 'DefaultLayout' extends keyof typeof __VLS_ctx ? typeof __VLS_ctx.DefaultLayout : (typeof __VLS_resolvedLocalAndGlobalComponents)['DefaultLayout'];
const __VLS_1 = __VLS_asFunctionalComponent(__VLS_0, new __VLS_0({ ...{}, }));
({} as { DefaultLayout: typeof __VLS_0; }).DefaultLayout;
({} as { DefaultLayout: typeof __VLS_0; }).DefaultLayout;
const __VLS_2 = __VLS_1({ ...{}, }, ...__VLS_functionalComponentArgsRest(__VLS_1));
const __VLS_3 = __VLS_pickFunctionalComponentCtx(__VLS_0, __VLS_2)!;
{
const __VLS_4 = ({} as __VLS_IntrinsicElements)["p"];
const __VLS_5 = __VLS_asFunctionalComponent(__VLS_4, {});
({} as __VLS_IntrinsicElements).p;
({} as __VLS_IntrinsicElements).p;
const __VLS_6 = __VLS_5({ ...{}, }, ...__VLS_functionalComponentArgsRest(__VLS_5));
const __VLS_7 = __VLS_pickFunctionalComponentCtx(__VLS_4, __VLS_6)!;
}
{
let __VLS_8!: 'Post' extends keyof typeof __VLS_ctx ? typeof __VLS_ctx.Post : (typeof __VLS_resolvedLocalAndGlobalComponents)['Post'];
const __VLS_9 = __VLS_asFunctionalComponent(__VLS_8, new __VLS_8({ ...{}, }));
({} as { Post: typeof __VLS_8; }).Post;
const __VLS_10 = __VLS_9({ ...{}, }, ...__VLS_functionalComponentArgsRest(__VLS_9));
const __VLS_11 = __VLS_pickFunctionalComponentCtx(__VLS_8, __VLS_10)!;
}
}
if (typeof __VLS_styleScopedClasses === 'object' && !Array.isArray(__VLS_styleScopedClasses)) {
}
var __VLS_slots!: {};
return __VLS_slots;
}

