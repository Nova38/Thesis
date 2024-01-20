// FILEPATH: /z:/source/repos/Thesis/pkg/web/utils/flatten.spec.ts
import { transformObject, type Mapping } from "./flatten";
import { describe, it, test, expect } from "vitest";

describe("transformObject", () => {
  it("should transform object keys based on mappings", () => {
    const obj = {
      name: "John Doe",
      age: 30,
      occupation: "Engineer",
    };

    const mappings: Mapping<
      typeof obj,
      { fullName: string; years: number; job: string }
    >[] = [
      { oldKey: "name", newKey: "fullName" },
      { oldKey: "age", newKey: "years" },
      { oldKey: "occupation", newKey: "job" },
    ];

    const result = transformObject(obj, mappings);
    expect(result).toEqual({
      fullName: "John Doe",
      years: 30,
      job: "Engineer",
    });
  });

  it("should apply transformation function if provided", () => {
    const obj = {
      name: "John Doe",
      age: "30",
      occupation: "Engineer",
    };

    const mappings: Mapping<
      typeof obj,
      { fullName: string; years: number; job: string }
    >[] = [
      { oldKey: "name", newKey: "fullName" },
      {
        oldKey: "age",
        newKey: "years",
        transform: (value: string) => parseInt(value),
      },
      { oldKey: "occupation", newKey: "job" },
    ];

    const result = transformObject(obj, mappings);
    expect(result).toEqual({
      fullName: "John Doe",
      years: 30,
      job: "Engineer",
    });
  });

  it("should apply general transformation function if provided", () => {
    const obj = {
      name: "John Doe",
      age: "30",
      occupation: "Engineer",
    };

    const mappings: Mapping<
      typeof obj,
      { fullName: string; years: number; job: string }
    >[] = [
      { oldKey: "name", newKey: "fullName" },
      { oldKey: "age", newKey: "years" },
      { oldKey: "occupation", newKey: "job" },
    ];

    const result = transformObject(obj, mappings, (value: string) =>
      value.toUpperCase(),
    );
    expect(result).toEqual({
      fullName: "JOHN DOE",
      years: "30",
      job: "ENGINEER",
    });
  });
});
