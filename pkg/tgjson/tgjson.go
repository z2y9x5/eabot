// Package tgjson.
// Telegram Bot API.
// Version 6.9 (2023-09-22).
// https://core.telegram.org/bots/api
package tgjson

// Ответ на запрос.
// Корневой элемент, содержащий все остальные элементы.
type Response struct {
	Ok          bool               `json:"ok"`
	Description string             `json:"description,omitempty"`
	Result      []Update           `json:"result,omitempty"`
	ErrorCode   int                `json:"error_code,omitempty"`
	Parameters  ResponseParameters `json:"parameters,omitempty"`
}

// Результат запроса.
// Доступен если в Response поле Ok == true.
type Update struct {
	UpdateId           int                `json:"update_id"`
	Message            Message            `json:"message,omitempty"`
	EditedMessage      Message            `json:"edited_message,omitempty"`
	ChannelPost        Message            `json:"channel_post,omitempty"`
	EditedChannelPost  Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      CallbackQuery      `json:"callback_query,omitempty"`
	ShippingQuery      ShippingQuery      `json:"shipping_query,omitempty"`
	PreCheckoutQuery   PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
	Poll               Poll               `json:"poll,omitempty"`
	PollAnswer         PollAnswer         `json:"poll_answer,omitempty"`
	MyChatMember       ChatMemberUpdated  `json:"my_chat_member,omitempty"`
	ChatMember         ChatMemberUpdated  `json:"chat_member,omitempty"`
	ChatJoinRequest    ChatJoinRequest    `json:"chat_join_request,omitempty"`
}

// Сообщение.
type Message struct {
	MessageId                     int                           `json:"message_id"`
	MessageThreadId               int                           `json:"message_thread_id,omitempty"`
	From                          User                          `json:"from,omitempty"`
	SenderChat                    Chat                          `json:"sender_chat,omitempty"`
	Date                          int                           `json:"date"`
	Chat                          Chat                          `json:"chat"`
	ForwardFrom                   User                          `json:"forward_from,omitempty"`
	ForwardFromChat               Chat                          `json:"forward_from_chat,omitempty"`
	ForwardFromMessageId          int                           `json:"forward_from_message_id,omitempty"`
	ForwardSignature              string                        `json:"forward_signature,omitempty"`
	ForwardSenderName             string                        `json:"forward_sender_name,omitempty"`
	ForwardDate                   int                           `json:"forward_date,omitempty"`
	IsTopicMessage                bool                          `json:"is_topic_message,omitempty"`
	IsAutomaticForward            bool                          `json:"is_automatic_forward,omitempty"`
	ReplyToMessage                *Message                      `json:"reply_to_message,omitempty"`
	ViaBot                        User                          `json:"via_bot,omitempty"`
	EditDate                      int                           `json:"edit_date,omitempty"`
	HasProtectedContent           bool                          `json:"has_protected_content,omitempty"`
	MediaGroupId                  string                        `json:"media_group_id,omitempty"`
	AuthorSignature               string                        `json:"author_signature,omitempty"`
	Text                          string                        `json:"text,omitempty"`
	Entities                      []MessageEntity               `json:"entities,omitempty"`
	Animation                     Animation                     `json:"animation,omitempty"`
	Audio                         Audio                         `json:"audio,omitempty"`
	Document                      Document                      `json:"document,omitempty"`
	Photo                         []PhotoSize                   `json:"photo,omitempty"`
	Sticker                       Sticker                       `json:"sticker,omitempty"`
	Story                         Story                         `json:"story,omitempty"`
	Video                         Video                         `json:"video,omitempty"`
	VideoNote                     VideoNote                     `json:"video_note,omitempty"`
	Voice                         Voice                         `json:"voice,omitempty"`
	Caption                       string                        `json:"caption,omitempty"`
	CaptionEntities               []MessageEntity               `json:"caption_entities,omitempty"`
	HasMediaSpoiler               bool                          `json:"has_media_spoiler,omitempty"`
	Contact                       Contact                       `json:"contact,omitempty"`
	Dice                          Dice                          `json:"dice,omitempty"`
	Game                          Game                          `json:"game,omitempty"`
	Poll                          Poll                          `json:"poll,omitempty"`
	Venue                         Venue                         `json:"venue,omitempty"`
	Location                      Location                      `json:"location,omitempty"`
	NewChatMembers                []User                        `json:"new_chat_members,omitempty"`
	LeftChatMember                User                          `json:"left_chat_member,omitempty"`
	NewChatTitle                  string                        `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []PhotoSize                   `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               bool                          `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              bool                          `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         bool                          `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            bool                          `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatId               int64                         `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId             int64                         `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *Message                      `json:"pinned_message,omitempty"`
	Invoice                       Invoice                       `json:"invoice,omitempty"`
	SuccessfulPayment             SuccessfulPayment             `json:"successful_payment,omitempty"`
	UserShared                    UserShared                    `json:"user_shared,omitempty"`
	ChatShared                    ChatShared                    `json:"chat_shared,omitempty"`
	ConnectedWebsite              string                        `json:"connected_website,omitempty"`
	WriteAccessAllowed            WriteAccessAllowed            `json:"write_access_allowed,omitempty"`
	PassportData                  PassportData                  `json:"passport_data,omitempty"`
	ProximityAlertTriggered       ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	ForumTopicCreated             ForumTopicCreated             `json:"forum_topic_created,omitempty"`
	ForumTopicEdited              ForumTopicEdited              `json:"forum_topic_edited,omitempty"`
	ForumTopicClosed              ForumTopicClosed              `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened            ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`
	GeneralForumTopicHidden       GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden     GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`
	VideoChatScheduled            VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted              VideoChatStarted              `json:"video_chat_started,omitempty"`
	VideoChatEnded                VideoChatEnded                `json:"video_chat_ended,omitempty"`
	VideoChatPartictipantsInvited VideoChatPartictipantsInvited `json:"video_chat_participants_invited,omitempty"`
	WebAppData                    WebAppData                    `json:"web_app_data,omitempty"`
	ReplyMarkup                   InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
}

// Пользователь Telegram или бот.
type User struct {
	Id                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name,omitempty"`
	Username                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	IsPremium               bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
}

// Чат.
type Chat struct {
	Id                                 int64           `json:"id"`
	Type                               string          `json:"type"`
	Title                              string          `json:"title,omitempty"`
	Username                           string          `json:"username,omitempty"`
	FirstName                          string          `json:"first_name,omitempty"`
	LastName                           string          `json:"last_name,omitempty"`
	IsForum                            bool            `json:"is_forum,omitempty"`
	Photo                              ChatPhoto       `json:"photo,omitempty"`
	ActiveUsernames                    []string        `json:"active_usernames,omitempty"`
	EmojiStatusCustomEmojiId           string          `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate          int             `json:"emoji_status_expiration_date,omitempty"`
	Bio                                string          `json:"bio,omitempty"`
	HasPrivateForwards                 bool            `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages bool            `json:"has_restricted_voice_and_video_messages,omitempty"`
	JoinToSendMessages                 bool            `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      bool            `json:"join_by_request,omitempty"`
	Description                        string          `json:"description,omitempty"`
	InviteLink                         string          `json:"invite_link,omitempty"`
	PinnedMessage                      *Message        `json:"pinned_message,omitempty"`
	Permissions                        ChatPermissions `json:"permissions,omitempty"`
	SlowModeDelay                      int             `json:"slow_mode_delay,omitempty"`
	MessageAutoDeleteTime              int             `json:"message_auto_delete_time,omitempty"`
	HasAggressiveAntiSpamEnabled       bool            `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers                   bool            `json:"has_hidden_members,omitempty"`
	HasProtectedContent                bool            `json:"has_protected_content,omitempty"`
	StickerSetName                     string          `json:"sticker_set_name,omitempty"`
	CanSetStickerSet                   bool            `json:"can_set_sticker_set,omitempty"`
	LinkedChatId                       int64           `json:"linked_chat_id,omitempty"`
	Location                           ChatLocation    `json:"location,omitempty"`
}

// Фото чата.
type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
}

// Действия, разрешенные пользователям в чате, без прав администратора.
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendAudios         bool `json:"can_send_audios,omitempty"`
	CanSendDocuments      bool `json:"can_send_documents,omitempty"`
	CanSendPhotos         bool `json:"can_send_photos,omitempty"`
	CanSendVideos         bool `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     bool `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
	CanManageTopics       bool `json:"can_manage_topics,omitempty"`
}

// Место, к которому привязан чат.
type ChatLocation struct {
	Location Location `json:"location"`
	Address  string   `json:"address"`
}

// Точка на карте.
type Location struct {
	Longitude            float64 `json:"longitude"`
	Latitude             float64 `json:"latitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

// Одиночный специальный элемент в текстовом сообщении.
// Например хэштеги, имена пользователей, URL-адреса.
type MessageEntity struct {
	Type          string `json:"type"`
	Offset        int    `json:"offset"`
	Length        int    `json:"length"`
	Url           string `json:"url,omitempty"`
	User          User   `json:"user,omitempty"`
	Language      string `json:"language,omitempty"`
	CustomEmojiId string `json:"custom_emoji_id,omitempty"`
}

// Анимированный файл (GIF или H.264/MPEG-4 AVC видео без звука).
type Animation struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	Duration     int       `json:"duration"`
	Thumbnail    PhotoSize `json:"thumbnail,omitempty"`
	FileName     string    `json:"file_name,omitempty"`
	MimeType     string    `json:"mime_type,omitempty"`
	FileSize     int64     `json:"file_size,omitempty"`
}

// Размер фото или миниатюры файла/стикера.
type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size,omitempty"`
}

// Аудиофайл, который клиент Telegram будет рассматривать как музыку.
type Audio struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Duration     int       `json:"duration"`
	Performer    string    `json:"performer,omitempty"`
	Title        string    `json:"title,omitempty"`
	FileName     string    `json:"file_name,omitempty"`
	MimeType     string    `json:"mime_type,omitempty"`
	FileSize     int64     `json:"file_size,omitempty"`
	Thumbnail    PhotoSize `json:"thumbnail,omitempty"`
}

// Обычный файл (в отличие от фотографий, голосовых сообщений и аудиофайлов).
type Document struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Thumbnail    PhotoSize `json:"thumbnail,omitempty"`
	FileName     string    `json:"file_name,omitempty"`
	MimeType     string    `json:"mime_type,omitempty"`
	FileSize     int64     `json:"file_size,omitempty"`
}

// Стикер.
type Sticker struct {
	FileId           string       `json:"file_id"`
	FileUniqueId     string       `json:"file_unique_id"`
	Type             string       `json:"type"`
	Width            int          `json:"width"`
	Height           int          `json:"height"`
	IsAnimated       bool         `json:"is_animated"`
	IsVideo          bool         `json:"is_video"`
	Thumbnail        PhotoSize    `json:"thumbnail,omitempty"`
	Emoji            string       `json:"emoji,omitempty"`
	SetName          string       `json:"set_name,omitempty"`
	PremiumAnimation File         `json:"premium_animation,omitempty"`
	MaskPosition     MaskPosition `json:"mask_position,omitempty"`
	CustomEmojiId    string       `json:"custom_emoji_id,omitempty"`
	NeedsRepainting  bool         `json:"needs_repainting,omitempty"`
	FileSize         int          `json:"file_size,omitempty"`
}

// Файл готовый к скачиванию.
// Файл может быть скачан по ссылке https://api.telegram.org/file/bot<token>/<file_path>.
// Ссылка гарантировано будет корректной как минимум 1 час.
// Когда время ссылки истекло, новая может быть запрошена вызовом getFile.
// Максимальный размер файла для скачивания 20MB.
type File struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}

// Позиция на лице, где по умолчанию должна быть размещена маска.
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

// История.
type Story struct{}

// Видео.
type Video struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	Duration     int       `json:"duration"`
	Thumbnail    PhotoSize `json:"thumbnail,omitempty"`
	FileName     string    `json:"file_name,omitempty"`
	MimeType     string    `json:"mime_type,omitempty"`
	FileSize     int64     `json:"file_size,omitempty"`
}

// Видеосообщение.
type VideoNote struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Length       int       `json:"length"`
	Duration     int       `json:"duration"`
	Thumbnail    PhotoSize `json:"thumbnail,omitempty"`
	FileSize     int       `json:"file_size,omitempty"`
}

// Голосовая заметка.
type Voice struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type,omitempty"`
	FileSize     int64  `json:"file_size,omitempty"`
}

// Телефонный контакт.
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserId      int64  `json:"user_id,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}

// Анимированный эмодзи отображающий случайное значение.
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

// Игра.
type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	Animation    Animation       `json:"animation,omitempty"`
}

// Информация о голосовании.
type Poll struct {
	Id                    string          `json:"id"`
	Question              string          `json:"question"`
	Options               []PollOption    `json:"options"`
	TotalVoterCount       int             `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  string          `json:"type"`
	AllowsMultipleAnswers bool            `json:"allow_multiple_answers"`
	CorrectOptionId       int             `json:"correct_option_id,omitempty"`
	Explanation           string          `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            int             `json:"open_period,omitempty"`
	CloseDate             int             `json:"close_date,omitempty"`
}

// Информация об одном из вариантов в голосовании.
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

// Место встречи.
type Venue struct {
	Location        Location `json:"location"`
	Title           string   `json:"title"`
	Address         string   `json:"address"`
	FoursquareId    string   `json:"foursquare_id,omitempty"`
	FoursquareType  string   `json:"foursquare_type,omitempty"`
	GooglePlaceId   string   `json:"google_place_id,omitempty"`
	GooglePlaceType string   `json:"google_place_type,omitempty"`
}

// Служебное сообщение об изменении настроек таймера автоудаления.
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// Базовая информация о выставленном счете.
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

// Содержит базовую информацию об успешном платеже.
type SuccessfulPayment struct {
	Currency                string    `json:"currency"`
	TotalAmount             int       `json:"total_amount"`
	InvoicePayload          string    `json:"invoice_payload"`
	ShippingOptionId        string    `json:"shipping_option_id,omitempty"`
	OrderInfo               OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeId string    `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId string    `json:"provider_payment_charge_id"`
}

// Информация о заказе.
type OrderInfo struct {
	Name            string          `json:"name,omitempty"`
	PhoneNumber     string          `json:"phone_number,omitempty"`
	Email           string          `json:"email,omitempty"`
	ShippingAddress ShippingAddress `json:"shipping_address,omitempty"`
}

// Адрес доставки.
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

// Cодержит информацию о пользователе, чей идентификатор был передан боту
// с помощью кнопки KeyboardButtonRequestUser.
type UserShared struct {
	RequestId int   `json:"request_id"`
	UserId    int64 `json:"user_id"`
}

// Cодержит информацию о чате, идентификатор которого был передан боту
// с помощью кнопки KeyboardButtonRequestChat.
type ChatShared struct {
	RequestId int   `json:"request_id"`
	ChatId    int64 `json:"chat_id"`
}

// Служебное сообщение о пользователе, разрешающем боту писать сообщения
// после добавления его в меню вложений, запуска веб-приложения по ссылке
// или принятия явного запроса от веб-приложения, отправленного методом requestWriteAccess.
type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request,omitempty"`
	WebAppName         string `json:"web_app_name,omitempty"`
	FromAttachmentMenu bool   `json:"from_attachment_menu,omitempty"`
}

// Данные Telegram Passport, которыми пользователь поделился с ботом.
type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials EncryptedCredentials       `json:"credentials"`
}

// Документы или другие элементы Telegram Passport,
// которыми пользователь поделился с ботом.
type EncryptedPassportElement struct {
	Type        string         `json:"type"`
	Data        string         `json:"data,omitempty"`
	PhoneNumber string         `json:"phone_number,omitempty"`
	Email       string         `json:"email,omitempty"`
	Files       []PassportFile `json:"files,omitempty"`
	FrontSide   PassportFile   `json:"front_side,omitempty"`
	ReverseSide PassportFile   `json:"reverse_side,omitempty"`
	Selfie      PassportFile   `json:"selfie,omitempty"`
	Translation []PassportFile `json:"translation,omitempty"`
	Hash        string         `json:"hash"`
}

// Файлы, загруженные для Telegram Passport.
type PassportFile struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int    `json:"file_date"`
}

// Данные, необходимые для расшифровки и аутентификации EncryptedPassportElement.
type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

// Cодержимое служебного сообщения, отправляемого всякий раз,
// когда пользователь в чате провоцирует оповещение о приближении,
// установленное другим пользователем.
type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"`
	Watcher  User `json:"watcher"`
	Distance int  `json:"distance"`
}

// Служебное сообщение о создании новой темы в чате.
type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

// Служебное сообщение о редактировании темы в чате.
type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

// Служебное сообщение о закрытии темы в чате.
// На данный момент не содержит информации.
type ForumTopicClosed struct{}

// Служебное сообщение о переоткрытии темы в чате.
// На данный момент не содержит информации.
type ForumTopicReopened struct{}

// Служебное сообщение о скрытии главной темы в чате.
// На данный момент не содержит информации.
type GeneralForumTopicHidden struct{}

// Служебное сообщение о доступности главной темы в чате.
// На данный момент не содержит информации.
type GeneralForumTopicUnhidden struct{}

// Служебное сообщение о запланированном в чате видеочате.
type VideoChatScheduled struct {
	StartDate int `json:"start_date"`
}

// Служебное сообщение о начале видеочата в чате.
// На данный момент не содержит информации.
type VideoChatStarted struct{}

// Служебное сообщение об окончании видеочата в чате.
type VideoChatEnded struct {
	Duration int `json:"duration"`
}

// Служебное сообщение о новых участниках, присоединившихся в видеочату.
type VideoChatPartictipantsInvited struct {
	Users []User `json:"users"`
}

// Данные, отправляемые из веб-приложения боту.
type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

// Встроенная клавиатура, которая появляется прямо рядом с сообщением, к которому она относится.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// Одиночная кнопка на встроенной клавиатуре.
type InlineKeyboardButton struct {
	Text                         string                      `json:"text"`
	Url                          string                      `json:"url,omitempty"`
	CallbackData                 string                      `json:"callback_data,omitempty"`
	WebApp                       WebAppInfo                  `json:"web_app,omitempty"`
	LoginUrl                     LoginUrl                    `json:"login_url,omitempty"`
	SwitchInlineQuery            string                      `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string                      `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat  SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CallbackGame                 CallbackGame                `json:"callback_game,omitempty"`
	Pay                          bool                        `json:"pay,omitempty"`
}

// Описание веб-приложения.
type WebAppInfo struct {
	Url string `json:"url"`
}

// Параметр кнопки встроенной клавиатуры,
// используемый для автоматической авторизации пользозвателя.
type LoginUrl struct {
	Url                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

// Inline кнопка, переключающая пользователя в inline режим в выбранном чате,
// с необязательным inline запросом по умолчанию.
type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query,omitempty"`
	AllowUserChats    bool   `json:"allow_user_chats,omitempty"`
	AllowBotChats     bool   `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   bool   `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool   `json:"allow_channel_chats,omitempty"`
}

// Заполнитель, в данный момент не содержит информации.
type CallbackGame struct{}

// Входящий inline запрос.
type InlineQuery struct {
	Id       string   `json:"id"`
	From     User     `json:"from"`
	Query    string   `json:"query"`
	Offset   string   `json:"offset"`
	ChatType string   `json:"chat_type,omitempty"`
	Location Location `json:"location,omitempty"`
}

// Результат inline запроса, который был выбран пользователем
// и отправлен собеседнику в чате.
type ChosenInlineResult struct {
	ResultId        string   `json:"result_id"`
	From            User     `json:"from"`
	Location        Location `json:"location,omitempty"`
	InlineMessageId string   `json:"inline_message_id,omitempty"`
	Query           string   `json:"query"`
}

// Входящий callback запрос от callback кнопки на встроенной клавиатуре.
type CallbackQuery struct {
	Id              string  `json:"id"`
	From            User    `json:"from"`
	Message         Message `json:"message,omitempty"`
	InlineMessageId string  `json:"inline_message_id,omitempty"`
	ChatInstance    string  `json:"chat_instance"`
	Data            string  `json:"data,omitempty"`
	GameShortName   string  `json:"game_short_name,omitempty"`
}

// Информация о поступившем запросе на доставку.
type ShippingQuery struct {
	Id              string          `json:"id"`
	From            User            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

// Информация о поступившем запросе на предзаказ.
type PreCheckoutQuery struct {
	Id               string    `json:"id"`
	From             User      `json:"from"`
	Currency         string    `json:"currency"`
	TotalAmount      int       `json:"total_amount"`
	InvoicePayload   string    `json:"invoice_payload"`
	ShippingOptionId string    `json:"shipping_option_id,omitempty"`
	OrderInfo        OrderInfo `json:"order_info,omitempty"`
}

// Ответ пользователя в неанонимном опросе.
type PollAnswer struct {
	PollId    string `json:"poll_id"`
	VoterChat Chat   `json:"voter_chat,omitempty"`
	User      User   `json:"user,omitempty"`
	OptionIds []int  `json:"option_ids"`
}

// Изменения в статусе участника чата.
type ChatMemberUpdated struct {
	Chat                    Chat           `json:"chat"`
	From                    User           `json:"from"`
	Date                    int            `json:"date"`
	OldChatMember           ChatMember     `json:"old_chat_member"`
	NewChatMember           ChatMember     `json:"new_chat_member"`
	InviteLink              ChatInviteLink `json:"invite_link,omitempty"`
	ViaChatFolderInviteLink bool           `json:"via_chat_folder_invite_link,omitempty"`
}

// Содержит информацию об отдельном участнике чата.
// На данный момент поддерживается 6 типов участников
// ChatMemberOwner
// ChatMemberAdministrator
// ChatMemberMember
// ChatMemberRestricted
// ChatMemberLeft
// ChatMemberBanned
type ChatMember struct {
	// Для всех участников
	Status string `json:"status"`
	User   User   `json:"user"`
	// Для ChatMemberBanned и ChatMemberRestricted
	UntilDate int `json:"until_date"`
	// Для ChatMemberOwner и ChatMemberAdministrator
	IsAnonymous bool   `json:"is_anonymous"`
	CustomTitle string `json:"custom_title,omitempty"`
	// Для ChatMemberAdministrator и ChatMemberRestricted
	CanChangeInfo   bool `json:"can_change_info"`
	CanInviteUsers  bool `json:"can_invite_users"`
	CanPinMessages  bool `json:"can_pin_messages"`
	CanManageTopics bool `json:"can_manage_topics"`
	// Для ChatMemberAdministrator
	CanBeEdited         bool `json:"can_be_edited"`
	CanManageChat       bool `json:"can_manage_chat"`
	CanDeleteMessages   bool `json:"can_delete_messages"`
	CanManageVideoChats bool `json:"can_manageVideoChats"`
	CanRestrictMembers  bool `json:"can_restrict_members"`
	CanPromoteMembers   bool `json:"can_promote_members"`
	CanPostMessages     bool `json:"can_post_messages,omitempty"`
	CanEditMessages     bool `json:"can_edit_messages,omitempty"`
	CanPostStories      bool `json:"can_post_stories,omitempty"`
	CanEditStories      bool `json:"can_edit_stories,omitempty"`
	CanDeleteStories    bool `json:"can_delete_stories,omitempty"`
	// Для ChatMemberRestricted
	IsMember              bool `json:"is_member"`
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendAudios         bool `json:"can_send_audios"`
	CanSendDocuments      bool `json:"can_send_documents"`
	CanSendPhotos         bool `json:"can_send_photos"`
	CanSendVideos         bool `json:"can_send_videos"`
	CanSendVideoNotes     bool `json:"can_send_video_notes"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
}

// Ссылка-приглашение в чат.
type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`
	Creator                 User   `json:"creator"`
	CreatesJoinRequest      bool   `json:"creates_join_request"`
	IsPrimary               bool   `json:"is_primary"`
	IsRevoked               bool   `json:"is_revoked"`
	Name                    string `json:"name,omitempty"`
	ExpireDate              int    `json:"expire_date,omitempty"`
	MemberLimit             int    `json:"member_limit,omitempty"`
	PendingJoinRequestCount int    `json:"pending_join_request_count,omitempty"`
}

// Запрос на присоединение, отправленный в чат.
type ChatJoinRequest struct {
	Chat       Chat           `json:"chat"`
	From       User           `json:"from"`
	UserChatId int64          `json:"user_chat_id"`
	Date       int            `json:"date"`
	Bio        string         `json:"bio,omitempty"`
	InviteLink ChatInviteLink `json:"invite_link,omitempty"`
}

// Описание причины неудачного запроса.
type ResponseParameters struct {
	MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int   `json:"retry_after,omitempty"`
}
