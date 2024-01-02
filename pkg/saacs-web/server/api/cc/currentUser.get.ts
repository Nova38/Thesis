import { useChaincode } from "~/server/utils/useChaincode";

export default defineEventHandler(async (event) => {
    const cc = await useChaincode(event);

    return cc.service.getCurrentUser();
});
