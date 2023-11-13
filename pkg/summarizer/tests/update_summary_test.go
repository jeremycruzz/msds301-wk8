package summarizer_test

import (
	"github.com/golang/mock/gomock"
	"github.com/jeremycruzz/msds301-wk8/pkg/mocks"
	"github.com/jeremycruzz/msds301-wk8/pkg/summarizer"
	"github.com/jeremycruzz/msds301-wk8/pkg/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UpdateSummary Unit Test", func() {
	var (
		mockCtrl          *gomock.Controller
		mockWikiApi       *mocks.MockWikipediaApi
		mockChatGpt       *mocks.MockChatGptService
		mockDb            *mocks.MockRepository
		summarizerService *summarizer.Service
		topic             string
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockWikiApi = mocks.NewMockWikipediaApi(mockCtrl)
		mockChatGpt = mocks.NewMockChatGptService(mockCtrl)
		mockDb = mocks.NewMockRepository(mockCtrl)
		topic = "anyTopic"
	})

	Describe("when the topic exists already", func() {

		BeforeEach(func() {
			mockDb.EXPECT().GetTitles().Return([]string{topic})
			mockDb.EXPECT().Get(topic).Return(&types.WikiTableInsert{PageID: 1, Title: topic, Extract: "MockExtract", Summary: "MockSummary", Eli5: "MockELI5"}, nil).Times(1)
			mockDb.EXPECT().Update(&types.WikiTableInsert{PageID: 1, Title: topic, Extract: "MockExtract", Summary: "NewMockSummary", Eli5: "NewMockELI5"}).Return(nil)
			mockChatGpt.EXPECT().Ask("MockExtract").Return("NewMockSummary", "NewMockELI5", nil)
			summarizerService = summarizer.New(mockWikiApi, mockChatGpt, mockDb)
		})

		AfterEach(func() {
			mockCtrl.Finish()
		})

		It("should return summary and ELI5 from DB", func() {
			// execute
			summary, eli5, err := summarizerService.UpdateSummary(topic)

			// assert
			Expect(err).To(BeNil())
			Expect(summary).To(Equal("NewMockSummary"))
			Expect(eli5).To(Equal("NewMockELI5"))
			Expect(summarizerService.ExistingTitles[topic]).To(BeTrue())
		})

		//TODO: Test error cases
	})

	Describe("when topic does not exist in DB", func() {
		BeforeEach(func() {
			mockDb.EXPECT().GetTitles().Return([]string{}) //no titles
			summarizerService = summarizer.New(mockWikiApi, mockChatGpt, mockDb)
		})
		Context("with no errors", func() {

			AfterEach(func() {
				mockCtrl.Finish()
			})

			It("should throw an error", func() {
				// execute
				summary, eli5, err := summarizerService.UpdateSummary(topic)

				// assert
				Expect(err).To(MatchError("topic does not exist"))
				Expect(summary).To(Equal(""))
				Expect(eli5).To(Equal(""))
				Expect(summarizerService.ExistingTitles[topic]).To(BeFalse())
			})
		})
	})
})
