import type { APIRoute } from "astro";
import { compatData, getBestCombos, getUsecase } from "../../../data/compat";

export function getStaticPaths() {
  return compatData.usecases.map((usecase) => ({
    params: { usecase: usecase.id },
  }));
}

export const GET: APIRoute = async ({ params }) => {
  const usecase = params.usecase ?? "";
  return new Response(
    JSON.stringify(
      {
        usecase,
        details: getUsecase(usecase),
        count: getBestCombos(usecase, 10).length,
        recommendations: getBestCombos(usecase, 10),
      },
      null,
      2,
    ),
    {
      headers: { "content-type": "application/json; charset=utf-8" },
    },
  );
};
