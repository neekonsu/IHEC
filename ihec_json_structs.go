package ihec

import (
	"errors"
	"fmt"
)

/*
Selection represents a collection of JSON metadata files for a selection of assays
the files entry is populated by the PopulateFiles() function in functions.go,
the Accessions entry is populated by a reflexive function locally in metadata.go
*/
type Selection struct {
	files      []METADATA
	Accessions map[string]bool
}

// LeanContext is a dense representation of a Dataset which only presents data relevant to the CEEHRC agreement
type LeanContext map[string]context

type context struct {
	Assay            string
	AssayCategory    string
	CellType         string
	CellTypeCategory string
	ReleasingGroup   string
	RawDataURL       string
	Accession        string
}

// METADATA represents one whole JSON object stored in the metadata for a single IHEC assay
type METADATA struct {
	Datasets map[string]struct {
		AnalysisAttributes struct {
			AlignmentSoftware        string `json:"alignment_software"`
			AlignmentSoftwareVersion string `json:"alignment_software_version"`
			AnalysisGroup            string `json:"analysis_group"`
			AnalysisSoftware         string `json:"analysis_software"`
			AnalysisSoftwareVersion  string `json:"analysis_software_version"`
		} `json:"analysis_attributes"`
		Browser map[string]struct {
			BigDataURL string `json:"big_data_url"`
			Md5Sum     string `json:"md5sum"`
			Primary    bool   `json:"primary"`
		} `json:"browser"`
		ExperimentAttributes struct {
			ExperimentType      string `json:"experiment_type"`
			ReferenceRegistryID string `json:"reference_registry_id"`
		} `json:"experiment_attributes"`
		IhecDataPortal struct {
			Assay            string `json:"assay"`
			AssayCategory    string `json:"assay_category"`
			CellType         string `json:"cell_type"`
			CellTypeCategory string `json:"cell_type_category"`
			ID               int    `json:"id"`
			PublishingGroup  string `json:"publishing_group"`
			RawDataURL       string `json:"raw_data_url"`
			ReleasingGroup   string `json:"releasing_group"`
		} `json:"ihec_data_portal"`
		OtherAttributes struct {
			BiomaterialProvider          string `json:"biomaterial_provider"`
			CollectionMethod             string `json:"collection_method"`
			DescriptionURL               string `json:"description_url"`
			DonorHealthStatusOntologyURI string `json:"donor_health_status_ontology_uri"`
			Gender                       string `json:"gender"`
			LibraryStrategy              string `json:"library_strategy"`
			Phenotype                    string `json:"phenotype"`
			Species                      string `json:"species"`
			SubjectID                    string `json:"subject_id"`
			TaxonID                      string `json:"taxon_id"`
		} `json:"other_attributes"`
		RawDataURL string `json:"raw_data_url"` /* This is what we want: METADATA.Datasets["KEY"].RawDataURL
		^^ Comes in following format: "https://www.ebi.ac.uk/ega/datasets/EGAD00001003963" */
		SampleID string `json:"sample_id"`
	} `json:"datasets"`
	HubDescription struct {
		Assembly        string `json:"assembly"`
		Date            string `json:"date"`
		Description     string `json:"description"`
		Email           string `json:"email"`
		PublishingGroup string `json:"publishing_group"`
		ReleasingGroup  string `json:"releasing_group"`
		TaxonID         int    `json:"taxon_id"`
	} `json:"hub_description"`
	Ok      bool `json:"ok"`
	Samples map[string]struct {
		BiomaterialType    string `json:"biomaterial_type"`
		Disease            string `json:"disease"`
		DiseaseOntologyURI string `json:"disease_ontology_uri"`
		DonorAge           int    `json:"donor_age"`
		DonorAgeUnit       string `json:"donor_age_unit"`
		DonorEthnicity     string `json:"donor_ethnicity"`
		DonorHealthStatus  string `json:"donor_health_status"`
		DonorID            string `json:"donor_id"`
		DonorLifeStage     string `json:"donor_life_stage"`
		DonorSex           string `json:"donor_sex"`
		SampleID           string `json:"sample_id"`
		SampleOntologyURI  string `json:"sample_ontology_uri"`
		TissueDepot        string `json:"tissue_depot"`
		TissueType         string `json:"tissue_type"`
	} `json:"samples"`
	Status int `json:"status"`
}

/*
PopulateAccessions iterates the Selection's files and populates Accessions with the appropiate strings
Accession comes in following format: "EGAD00001003963"
Because I changed the Accessions datatype to map[string]bool, there are no duplicate items.
*/
func (s *Selection) PopulateAccessions() {
	output := make(map[string]bool)
	for _, metadata := range s.files {
		for _, dataset := range metadata.Datasets {
			output[IsolateAccession(dataset.RawDataURL)] = false
			output[IsolateAccession(dataset.IhecDataPortal.RawDataURL)] = false
		}
	}
	delete(output, "")
	s.Accessions = output
}

// MakeLeanContext reflexively generates a LeanContext from a Selection; Selection.files is only required in order to complete
func (s *Selection) MakeLeanContext() (LeanContext, error) {
	var output LeanContext
	tmp := make(map[string]context)
	if len(s.files) == 0 {
		return nil, errors.New("provided Selection has empty files.\nAt least len(1) Selection.files is required")
	}
	for _, json := range s.files {
		for name, contents := range json.Datasets {
			// I can't figure out how I could have made LeanContext type map[string]struct{...};
			// I don't know how to initialize the anonymouus struct when dealing with a map.
			// Something to research for later
			tmp[name] = context{
				contents.IhecDataPortal.Assay,
				contents.IhecDataPortal.AssayCategory,
				contents.IhecDataPortal.CellType,
				contents.IhecDataPortal.CellTypeCategory,
				contents.IhecDataPortal.ReleasingGroup,
				contents.IhecDataPortal.RawDataURL,
				IsolateAccession(contents.IhecDataPortal.RawDataURL),
			}
		}
	}
	output = tmp
	return output, nil
}

// PrintAccessions pretty prints the unique accessions list along with the total number of unique accessions
func (s *Selection) PrintAccessions() {
	i := 0
	for key := range s.Accessions {
		fmt.Println(key)
		i++
	}
	fmt.Println("Total of", i, "unique accessions found")
}

// PrintLeanContext calls MakeLeanContext() and then pretty prints the results
func (s *Selection) PrintLeanContext() {
	lc, err := s.MakeLeanContext()
	CheckErr("Unable to PrintLeanContext: ", err)
	for key, val := range lc {
		fmt.Println(key + ":")
		fmt.Println("\t"+"Raw Data URl:", val.RawDataURL)
		fmt.Println("\t"+"Accession:", val.Accession)
		fmt.Println("\t"+"Assay:", val.Assay)
		fmt.Println("\t"+"Assay Category:", val.AssayCategory)
		fmt.Println("\t"+"Cell Type:", val.CellType)
		fmt.Println("\t"+"Cell Type Category:", val.CellTypeCategory)
		fmt.Println("\t"+"Releasing Group:", val.ReleasingGroup)
	}
}
