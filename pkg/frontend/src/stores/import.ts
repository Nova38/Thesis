import { defineStore, acceptHMRUpdate } from 'pinia';
// import { Specimen, SpecimenKeys, SpecimenMapping } from 'src/lib/specimen';
import { computed, reactive, ref, watch } from 'vue';
import type { Ref } from 'vue';
import Papa, { ParseResult } from 'papaparse';
import { api } from 'boot/axios';

export type ImportStatus = 'new' | 'uploading' | 'error' | 'success';

export interface MappedSpecimen extends Specimen {
  index: string;
  status: ImportStatus;
  statusMessage: string;
}

export const useImportStore = defineStore('csv_import', () => {
  const headers: Ref<string[]> = ref([]);
  const data = ref<Record<string, string>[]>();

  const specimenKeys = Object.keys(new Specimen()) as SpecimenKeys[];
  const specimenMapping = reactive(new Specimen());

  const selected = ref<Record<string, string>[]>([]);
  const visibleColumns = ref<string[] | null>(null);

  const collection = 'KU Ornithology';

  const file = ref<File | null>(null);

  const mappedColl = [
    {
      name: 'catalogNumber',
      field: 'catalogNumber',
      label: 'catalogNumber',
      sortable: true,
    },
    {
      name: 'accessionNumber',
      field: 'accessionNumber',
      label: 'accessionNumber',
      sortable: true,
    },
    {
      name: 'catalogDate',
      field: 'catalogDate',
      label: 'catalogDate',
      sortable: true,
    },
    {
      name: 'cataloger',
      field: 'cataloger',
      label: 'cataloger',
      sortable: true,
    },
    {
      name: 'determiner',
      field: 'determiner',
      label: 'determiner',
      sortable: true,
    },
    {
      name: 'determineDate',
      field: 'determineDate',
      label: 'determineDate',
      sortable: true,
    },
    {
      name: 'fieldNumber',
      field: 'fieldNumber',
      label: 'fieldNumber',
      sortable: true,
    },
    {
      name: 'fieldDate',
      field: 'fieldDate',
      label: 'fieldDate',
      sortable: true,
    },
    {
      name: 'collector',
      field: 'collector',
      label: 'collector',
      sortable: true,
    },
    {
      name: 'taxon_kingdom',
      field: 'taxon_kingdom',
      label: 'taxon_kingdom',
      sortable: true,
    },
    {
      name: 'taxon_phylum',
      field: 'taxon_phylum',
      label: 'taxon_phylum',
      sortable: true,
    },
    {
      name: 'taxon_class',
      field: 'taxon_class',
      label: 'taxon_class',
      sortable: true,
    },
    {
      name: 'taxon_order',
      field: 'taxon_order',
      label: 'taxon_order',
      sortable: true,
    },
    {
      name: 'taxon_family',
      field: 'taxon_family',
      label: 'taxon_family',
      sortable: true,
    },
    {
      name: 'taxon_genus',
      field: 'taxon_genus',
      label: 'taxon_genus',
      sortable: true,
    },
    {
      name: 'taxon_species',
      field: 'taxon_species',
      label: 'taxon_species',
      sortable: true,
    },
    { name: 'country', field: 'country', label: 'country', sortable: true },
    {
      name: 'stateProvince',
      field: 'stateProvince',
      label: 'stateProvince',
      sortable: true,
    },
    { name: 'county', field: 'county', label: 'county', sortable: true },
    { name: 'locality', field: 'locality', label: 'locality', sortable: true },
    { name: 'latitude', field: 'latitude', label: 'latitude', sortable: true },
    {
      name: 'longitude',
      field: 'longitude',
      label: 'longitude',
      sortable: true,
    },
    { name: 'habitat', field: 'habitat', label: 'habitat', sortable: true },
    {
      name: 'preparation',
      field: 'preparation',
      label: 'preparation',
      sortable: true,
    },
    {
      name: 'condition',
      field: 'condition',
      label: 'condition',
      sortable: true,
    },
    { name: 'notes', field: 'notes', label: 'notes', sortable: true },
    { name: 'image', field: 'image', label: 'image', sortable: true },
    { name: 'loans', field: 'loans', label: 'loans', sortable: true },
    { name: 'grants', field: 'grants', label: 'grants', sortable: true },
  ];

  const mappedArray = computed(() => {
    if (!data.value || !specimenMapping) {
      return;
    }

    const map = new Map(Object.entries(specimenMapping)) as SpecimenMapping;

    return selected.value.map((obj) => {
      const specimen = new Specimen() as MappedSpecimen;
      specimen.index = obj.index;
      specimen.status = obj.status as ImportStatus;

      map.forEach((value, key) => {
        if (key === 'vandalizedTransactions') {
          specimen[key] = [];
          return;
        } else if (key === 'docType') {
          return;
        } else {
          specimen[key] = obj[value] || '';
        }
      });
      return specimen;
    });
  });

  watch(file, (file) => {
    if (!file) {
      return;
    }

    Papa.parse(file, {
      // worker: true,
      header: true,
      complete: (results: ParseResult<Record<string, string>>) => {
        console.log(results);
        console.log(typeof results.data[0]);

        data.value = results.data.map((row) => {
          return Object.assign({ status: 'new' }, row);
        });

        visibleColumns.value = results.meta.fields || [];
        visibleColumns.value.unshift('status');

        headers.value = results.meta.fields || [];

        // if the headers match the Specimen class explicenly, then set the inital mapping to map those fields to the headers

        for (const key of specimenKeys) {
          if (headers.value.includes(key)) {
            if (key === 'vandalizedTransactions') {
              continue;
            }

            specimenMapping[key] = key;
          }
        }

        // add a column for the index
        headers.value.unshift('status');

        for (let i = 0; i < data.value.length; i++) {
          data.value[i].index = `${i}`;
          data.value[i].status = 'new';
        }
      },
    });
  });

  const clearKey = (key: SpecimenKeys) => {
    if (key === 'vandalizedTransactions') {
      specimenMapping[key] = [];
      return;
    } else if (key === 'docType') {
      return;
    } else {
      specimenMapping[key] = '';
    }
  };

  const upload = async () => {
    console.log('uploading');
    console.table(mappedArray.value);
    if (data.value) {
      const unUploaded = mappedArray.value?.filter(
        (specimen) =>
          data.value &&
          !(data.value[Number(specimen.index)].status === 'success')
      );

      if (unUploaded && unUploaded.length > 0) {
        unUploaded.forEach((specimen: MappedSpecimen) => {
          specimen.collection = collection;
        });

        const results = await api.post(
          'biochain/specimens-bulk/bulk',
          unUploaded
        );
      }
    }
    mappedArray.value?.forEach((specimen: MappedSpecimen) => {
      if (
        data.value &&
        data.value[Number(specimen.index)].status === 'success'
      ) {
        return;
      }

      api
        .post('biochain/specimens', specimen)
        .then((res) => {
          console.log(res);
          if (data.value) {
            // data.value[specimen.index as string].status = 'success';
            data.value[Number(specimen.index)].status = 'success';
          }
        })
        .catch((err) => {
          console.error(err);
          if (data.value) {
            // data.value[specimen.index as string].status = 'error';
            data.value[Number(specimen.index)].status = 'error';
            data.value[Number(specimen.index)].statusMessage = err;
          }
        });
    });
  };

  return {
    clearKey,
    upload,
    file,
    headers,
    data,
    selected,
    specimenMapping,
    specimenKeys,
    mappedArray,
    visibleColumns,
    mappedColl,
  };
});

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useImportStore, import.meta.hot));
}
