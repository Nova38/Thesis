<script lang="ts" setup>
import { schema } from 'src/lib/ccbio';

// import { Specimen as ZSpecimen } from 'src/lib/zod/state';
// import { DefaultFormProps } from 'src/composables/defaultProps';
import { useVModel } from '@vueuse/core';
import { PlainMessage } from '@bufbuild/protobuf';
import { FormSpecimen } from 'src/lib/ccbio/utils/specimen';

export interface Props {
  editPermission?: PlainMessage<schema.Permissions>;
  specimen: FormSpecimen;
  enableEdit?: boolean;
  showUpdater?: boolean;
  showHeader?: boolean;
}

// Compute wither or not the specimen is editable
const editable = computed((): PlainMessage<schema.Permissions> => {
  const allDisabled = new schema.Permissions({
    roles: false,
    users: false,
    specimen: false,
    primary: false,
    secondary: false,
    taxon: false,
    georeference: false,
    images: false,
    loans: false,
    grants: false,
    hidden: false,
  });

  const allEnabled = new schema.Permissions({
    roles: true,
    users: true,
    specimen: true,
    primary: true,
    secondary: true,
    taxon: true,
    georeference: true,
    images: true,
    loans: true,
    grants: true,
    hidden: true,
  });

  if (!props.enableEdit) return allDisabled;
  if (props.enableEdit) return allEnabled;

  // if (!props.editPermission) return allEnabled;

  // and all of the properties from props.editPermission to props.enableEdit
  return allEnabled;
});

const props = withDefaults(defineProps<Props>(), {
  editable: false,
  enableEdit: true,
  showUpdater: true,
  showHeader: true,
  showEditOptions: true,
});
const emit = defineEmits(['update:specimen']);

const specimen = useVModel(props, 'specimen', emit);
//    ^?

const dateMask = 'M/D/YYYY';

// const getJson = (obj: schema.Specimen) => {
//   return schema.Specimen.toJSON(obj);
// };

const DefaultFormProps = ref({
  outlined: true,
  dense: true,
  'stack-label': true,
  // standout: true,
  // rounded: true,
  // filled: true,
});
</script>

<template>
  <q-card id="wrapper" v-if="specimen">
    <slot name="HeaderBar" />
    <q-card-section
      v-if="props.showHeader"
      class="bg-secondary text-white"
      id="HeaderWrapper"
    >
      <div class="row">
        <div v-if="specimen?.taxon" class="text-h5">
          {{ specimen.taxon?.genus + ' ' + specimen.taxon?.species }}
        </div>
        <div
          v-if="specimen?.primary"
          class="text-subtitle2"
          style="margin-left: auto"
        >
          <q-chip
            color="accent"
            square
            class="chip"
            v-if="specimen.primary.catalogNumber"
            :label="'Catalog Number: ' + specimen.primary?.catalogNumber"
          />
          <q-chip
            color="accent"
            class="chip"
            square
            v-if="specimen.primary.tissueNumber"
            :label="'Tissue Number: ' + specimen.primary.tissueNumber"
          />
        </div>
      </div>
    </q-card-section>
    <q-form>
      <q-expansion-item
        expand-icon-toggle
        default-opened
        :icon="!editable.taxon ? 'ti-lock' : 'ti-pencil-alt'"
        label="Taxon"
        header-class="text-center section-header"
      >
        <q-card>
          <q-card-section v-if="specimen?.taxon" class="">
            <div class="row justify-evenly q-mb-sm">
              <q-input
                v-model="specimen.taxon.kingdom"
                class="col-auto q-mt-sm q-mb-sm"
                label="Kingdom"
                v-bind="DefaultFormProps"
                :disable="!editable.taxon"
              />
              <q-input
                v-model="specimen.taxon.phylum"
                class="col-auto q-mt-sm q-mb-sm"
                label="Phylum"
                v-bind="DefaultFormProps"
                :disable="!editable.taxon"
              />
              <q-input
                v-model="specimen.taxon.order"
                class="col-auto q-mt-sm q-mb-sm"
                label="Order"
                v-bind="DefaultFormProps"
                :disable="!editable.taxon"
              />
              <q-input
                v-model="specimen.taxon.family"
                class="col-auto q-mt-sm q-mb-sm"
                label="Family"
                v-bind="DefaultFormProps"
                :disable="!editable.taxon"
              />
              <q-input
                v-model="specimen.taxon.genus"
                class="col-auto q-mt-sm q-mb-sm"
                label="Genus"
                v-bind="DefaultFormProps"
                :disable="!editable.taxon"
              />
              <q-input
                v-model="specimen.taxon.species"
                class="col-auto q-mt-sm q-mb-sm"
                label="Species"
                v-bind="DefaultFormProps"
                :disable="!editable.taxon"
              />
              <q-input
                v-model="specimen.taxon.subspecies"
                class="col-auto q-mt-sm q-mb-sm"
                label="Subspecies"
                v-bind="DefaultFormProps"
                :disable="!editable.taxon"
              />
            </div>
            <q-item-label overline><b>Taxon Last Updater</b></q-item-label>
            <q-item-label class="text-bold">
              {{ specimen.taxon.lastModified }}
            </q-item-label>
          </q-card-section>
        </q-card>
      </q-expansion-item>

      <q-separator inset />

      <q-expansion-item
        expand-icon-toggle
        default-opened
        :icon="!editable.primary ? 'ti-lock' : 'ti-pencil-alt'"
        label="Primary"
        header-class="text-center section-header"
      >
        <q-card>
          <div v-if="specimen?.primary">
            <q-card-section class="q-mt-sm q-mb-sm">
              <!-- <div class="text-h6">Overline</div> -->
              <div class="row justify-evenly q-mb-sm">
                <q-input
                  v-model="specimen.primary.tissueNumber"
                  label="Tissue Number"
                  class="col q-ma-sm"
                  v-bind="DefaultFormProps"
                  :disable="!editable.primary"
                />
                <q-input
                  v-model="specimen.primary.accessionNumber"
                  class="col q-ma-sm"
                  label="Accession Number"
                  v-bind="DefaultFormProps"
                  :disable="!editable.primary"
                />
              </div>
              <div class="row justify-evenly q-mb-sm">
                <q-input
                  v-model="specimen.primary.cataloger"
                  class="col q-ma-sm"
                  label="Cataloger"
                  v-bind="DefaultFormProps"
                  :disable="!editable.primary"
                />
                <q-input
                  v-model="specimen.primary.catalogNumber"
                  class="col q-ma-sm"
                  label="Catalog Number"
                  v-bind="DefaultFormProps"
                  :disable="!editable.primary"
                />
                <div>
                  <q-input
                    v-model="specimen.primary.catalogDate"
                    class="col q-ma-sm"
                    label="Catalog Date"
                    v-bind="DefaultFormProps"
                    :disable="!editable.primary"
                  >
                    <template #append>
                      <q-icon name="event" class="cursor-pointer">
                        <q-popup-proxy
                          cover
                          transition-show="scale"
                          transition-hide="scale"
                        >
                          <q-date
                            :mask="dateMask"
                            v-model="specimen.primary.catalogDate"
                          >
                            <div class="row items-center justify-end">
                              <q-btn
                                v-close-popup
                                label="Close"
                                color="primary"
                                flat
                              />
                            </div>
                          </q-date>
                        </q-popup-proxy>
                      </q-icon>
                    </template>
                  </q-input>
                </div>
              </div>
              <div class="row justify-evenly q-mb-sm">
                <q-input
                  v-model="specimen.primary.collector"
                  label="Collector"
                  class="col q-ma-sm"
                  v-bind="DefaultFormProps"
                  :disable="!editable.primary"
                />
                <q-input
                  v-model="specimen.primary.fieldNumber"
                  class="col q-ma-sm"
                  label="Field Number"
                  v-bind="DefaultFormProps"
                  :disable="!editable.primary"
                />
                <q-input
                  date
                  v-model="specimen.primary.fieldDate"
                  class="col q-ma-sm"
                  label="Field Date"
                  v-bind="DefaultFormProps"
                  :disable="!editable.primary"
                >
                  <template #append>
                    <q-icon name="event" class="cursor-pointer">
                      <q-popup-proxy
                        cover
                        transition-show="scale"
                        transition-hide="scale"
                      >
                        <q-date
                          :mask="dateMask"
                          v-model="specimen.primary.fieldDate"
                        >
                          <div class="row items-center justify-end">
                            <q-btn
                              v-close-popup
                              label="Close"
                              color="primary"
                              flat
                            />
                          </div>
                        </q-date>
                      </q-popup-proxy>
                    </q-icon>
                  </template>
                </q-input>
              </div>

              <div class="row justify-evenly q-mb-sm">
                <q-input
                  v-model="specimen.primary.determiner"
                  class="col q-ma-sm"
                  label="Determiner"
                  v-bind="DefaultFormProps"
                  :disable="!editable.primary"
                />
                <q-input
                  v-model="specimen.primary.determinedReason"
                  label="Determined Reason"
                  class="col q-ma-sm"
                  v-bind="DefaultFormProps"
                  :disable="!editable.primary"
                />
                <q-input
                  v-model="specimen.primary.determinedDate"
                  class="col q-ma-sm"
                  label="Determine Date"
                  v-bind="DefaultFormProps"
                  :disable="!editable.primary"
                >
                  <template #append>
                    <q-icon name="event" class="cursor-pointer">
                      <q-popup-proxy
                        cover
                        transition-show="scale"
                        transition-hide="scale"
                      >
                        <q-date
                          v-model="specimen.primary.determinedDate"
                          :mask="dateMask"
                        >
                          <div class="row items-center justify-end">
                            <q-btn
                              v-close-popup
                              label="Close"
                              color="primary"
                              flat
                            />
                          </div>
                        </q-date>
                      </q-popup-proxy>
                    </q-icon>
                  </template>
                </q-input>
              </div>
            </q-card-section>
          </div>
          <q-card-section v-if="specimen.id">
            <q-input
              v-model="specimen.id.collectionId"
              label="Collection"
              class="q-ma-sm q-mb-sm"
              v-bind="DefaultFormProps"
              :disable="true"
            />

            <q-item-label overline><b>Primary Last Updater</b></q-item-label>
            <q-item-label class="text-bold">
              {{ specimen.primary.lastModified }}
            </q-item-label>
          </q-card-section>
        </q-card>
      </q-expansion-item>

      <q-separator inset />

      <q-expansion-item
        expand-icon-toggle
        default-opened
        :icon="!editable.georeference ? 'ti-lock' : 'ti-pencil-alt'"
        label="Georeference"
        header-class="text-center"
      >
        <q-card>
          <q-card-section v-if="specimen?.georeference" class="q-mt-sm q-mb-sm">
            <div class="row justify-evenly q-mb-sm">
              <q-input
                v-model="specimen.georeference.country"
                label="Country"
                class="col q-ma-sm"
                v-bind="DefaultFormProps"
                :disable="!editable.georeference"
              />
              <q-input
                v-model="specimen.georeference.stateProvince"
                label="State/Province"
                class="col q-ma-sm"
                v-bind="DefaultFormProps"
                :disable="!editable.georeference"
              />
              <q-input
                v-model="specimen.georeference.county"
                label="County"
                class="col q-ma-sm"
                v-bind="DefaultFormProps"
                :disable="!editable.georeference"
              />
            </div>

            <div class="row justify-evenly q-mb-sm">
              <q-input
                v-model="specimen.georeference.habitat"
                class="col q-ma-sm"
                label="Habitat"
                type="textarea"
                autogrow
                v-bind="DefaultFormProps"
                :disable="!editable.georeference"
              />
              <q-input
                v-model="specimen.georeference.locality"
                class="col q-ma-sm"
                label="Locality"
                v-bind="DefaultFormProps"
                :disable="!editable.georeference"
              />
            </div>

            <div class="row justify-evenly q-mb-sm">
              <q-input
                v-model="specimen.georeference.longitude"
                class="col q-ma-sm"
                label="Longitude"
                v-bind="DefaultFormProps"
                :disable="!editable.georeference"
              />

              <q-input
                v-model="specimen.georeference.latitude"
                class="col q-ma-sm"
                label="Latitude"
                v-bind="DefaultFormProps"
                :disable="!editable.georeference"
              />
            </div>
            <q-item-label overline
              ><b>Georeference Last Updater</b></q-item-label
            >
            <q-item-label class="text-bold">
              {{ specimen.secondary.lastModified }}
            </q-item-label>
          </q-card-section>
        </q-card>
      </q-expansion-item>

      <q-separator inset />

      <q-expansion-item
        expand-icon-toggle
        default-opened
        :icon="!editable.secondary ? 'ti-lock' : 'ti-pencil-alt'"
        label="Secondary"
        header-class="text-center section-header"
      >
        <q-card>
          <q-card-section v-if="specimen?.secondary" class="q-mt-sm q-mb-sm">
            <q-input
              v-model="specimen.secondary.condition"
              class="col-auto q-mt-sm q-mb-sm"
              label="Condition"
              type="textarea"
              autogrow
              v-bind="DefaultFormProps"
              :disable="!editable.secondary"
            />

            <q-input
              v-model="specimen.secondary.preparation"
              class="col-auto q-mt-sm q-mb-sm"
              label="Preparation"
              type="textarea"
              autogrow
              v-bind="DefaultFormProps"
              :disable="!editable.secondary"
            />
            <q-input
              v-model="specimen.secondary.notes"
              class="col-auto q-mt-sm q-mb-sm"
              label="Notes"
              type="textarea"
              autogrow
              v-bind="DefaultFormProps"
              :disable="!editable.secondary"
            />
            <q-item-label overline><b>Secondary Last Updater</b></q-item-label>
            <q-item-label class="text-bold">
              {{ specimen.secondary.lastModified }}
            </q-item-label>
          </q-card-section>
        </q-card>
      </q-expansion-item>

      <q-separator inset />

      <q-expansion-item
        expand-icon-toggle
        default-opened
        :icon="!editable.loans ? 'ti-lock' : 'ti-pencil-alt'"
        label="Loans"
        header-class="text-center section-header"
      >
        <q-card>
          <q-card-section class="">
            <q-input
              v-model="specimen.loans"
              class="col-auto q-mt-sm q-mb-sm"
              label="Loans"
              type="textarea"
              autogrow
              v-bind="DefaultFormProps"
              :disable="!editable.loans"
            />
          </q-card-section>
        </q-card>
      </q-expansion-item>

      <q-separator inset />

      <q-expansion-item
        expand-icon-toggle
        default-opened
        :icon="!editable.grants ? 'ti-lock' : 'ti-pencil-alt'"
        label="Grants"
        header-class="text-center section-header"
      >
        <q-card>
          <q-card-section class="">
            <q-input
              v-model="specimen.grants"
              class="col-auto q-mt-sm q-mb-sm"
              label="Grants"
              type="textarea"
              autogrow
              v-bind="DefaultFormProps"
              :disable="!editable.grants"
            />
          </q-card-section>
        </q-card>
      </q-expansion-item>
      <q-separator inset />

      <q-card-section v-if="showUpdater">
        <q-item-label overline><b>Last Updater</b></q-item-label>
        <q-item-label class="text-bold">
          {{ specimen.lastModified }}
        </q-item-label>
      </q-card-section>
      <q-separator inset />

      <q-card-section>
        <q-card-actions>
          <slot name="actions" />
        </q-card-actions>
      </q-card-section>
    </q-form>
  </q-card>
  <!-- <q-btn @click="toggleEdit"> hi </q-btn> -->
</template>

<style scoped>
.chip {
  color: #fff;
  justify-content: flex-end;
  /* align-self: flex-end; */
}

.q-field__label {
  color: #08b865;
  font-weight: bold;
  display: block;
}
</style>
