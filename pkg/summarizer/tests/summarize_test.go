package summarizer_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	"github.com/jeremycruzz/msds301-wk8/pkg/mocks"
	"github.com/jeremycruzz/msds301-wk8/pkg/summarizer"
	"github.com/jeremycruzz/msds301-wk8/pkg/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Summarize Unit Test", func() {
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
			mockDb.EXPECT().Get(topic).Return(&types.WikiTableInsert{Summary: "MockSummary", Eli5: "MockELI5"}, nil)
			summarizerService = summarizer.New(mockWikiApi, mockChatGpt, mockDb)
		})

		AfterEach(func() {
			mockCtrl.Finish()
		})

		It("should return summary and ELI5 from DB", func() {
			// execute
			summary, eli5, err := summarizerService.Summarize(topic)

			// assert
			Expect(err).To(BeNil())
			Expect(summary).To(Equal("MockSummary"))
			Expect(eli5).To(Equal("MockELI5"))
			Expect(summarizerService.ExistingTitles[topic]).To(BeTrue())
		})
	})

	Describe("when topic does not exist in DB", func() {
		BeforeEach(func() {
			mockDb.EXPECT().GetTitles().Return([]string{}) //no titles
			summarizerService = summarizer.New(mockWikiApi, mockChatGpt, mockDb)
		})
		Context("with no errors", func() {
			BeforeEach(func() {
				mockWikiApi.EXPECT().Query(topic).Return(&types.WikiPage{Extract: "MockExtract"}, nil)
				mockChatGpt.EXPECT().Ask("MockExtract").Return("MockSummary", "MockELI5", nil)
				mockDb.EXPECT().Insert(gomock.Any()).Return(nil)
			})

			AfterEach(func() {
				mockCtrl.Finish()
			})

			It("should query Wiki API and ChatGPT, store in DB, then add to existing titles", func() {
				// execute
				summary, eli5, err := summarizerService.Summarize(topic)

				// assert
				Expect(err).To(BeNil())
				Expect(summary).To(Equal("MockSummary"))
				Expect(eli5).To(Equal("MockELI5"))
				Expect(summarizerService.ExistingTitles[topic]).To(BeTrue())
			})
		})

		Context("with wikiapi errors", func() {

			BeforeEach(func() {
				mockWikiApi.EXPECT().Query(topic).Return(nil, errors.New("wiki api error"))
			})

			AfterEach(func() {
				mockCtrl.Finish()
			})

			It("should handle errors from Wiki API", func() {
				// execute
				_, _, err := summarizerService.Summarize(topic)

				// assert
				Expect(err).To(MatchError("wiki api error"))
			})
		})

		Context("with chatgpt errors", func() {

			BeforeEach(func() {
				mockWikiApi.EXPECT().Query(topic).Return(&types.WikiPage{Extract: "MockExtract"}, nil)
				mockChatGpt.EXPECT().Ask("MockExtract").Return("", "", errors.New("chatgpt error"))
			})

			AfterEach(func() {
				mockCtrl.Finish()
			})

			It("should handle errors from chatgpt", func() {
				// execute
				_, _, err := summarizerService.Summarize(topic)

				// assert
				Expect(err).To(MatchError("chatgpt error"))
			})
		})

		Context("with db errors", func() {

			BeforeEach(func() {
				mockWikiApi.EXPECT().Query(topic).Return(&types.WikiPage{Extract: "MockExtract"}, nil)
				mockChatGpt.EXPECT().Ask("MockExtract").Return("MockSummary", "MockELI5", nil)
				mockDb.EXPECT().Insert(gomock.Any()).Return(errors.New("db insert error")) //TODO: Tighten up parameter
			})

			AfterEach(func() {
				mockCtrl.Finish()
			})

			It("should handle errors from db", func() {
				// execute
				_, _, err := summarizerService.Summarize(topic)

				// assert
				Expect(err).To(MatchError("db insert error"))
			})
		})
	})
})
