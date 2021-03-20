package diploma

type Service struct {
	logger Logger
}

func New(logger Logger) *Service {
	return &Service{logger: logger}
}

func (s *Service) GeneratePDF(id string, data interface{}) error {
	body, err := parseTemplate(data)
	if err != nil {
		return err
	}

	err = generatePDF(id, body)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GeneratePDFNoHTML(id string, data interface{}) error {
	// body, err := parseTemplate(data)
	// if err != nil {
	// 	return err
	// }

	err := generatePDFNoHTML(id, "")
	if err != nil {
		return err
	}

	return nil
}
