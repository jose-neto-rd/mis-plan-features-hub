package constants

const (
	// Plans
	BasicPlan    = "basic"
	ProPlan      = "pro"
	AdvancedPlan = "advanced"

	// Chat module features
	ChatChat                 = "chat:chat"                   // Chat de atendimento
	ChatQuickMessages        = "chat:quick-messages"         // Mensagens rápidas e variáveis
	ChatConversationID       = "chat:conversation-id"        // Identificação de cada atendimento (ID da conversa)
	ChatRealTimeMonitoring   = "chat:realtime-monitoring"    // Monitoramento em tempo real da equipe
	ChatRating               = "chat:rating"                 // Avaliação de atendimento
	ChatConversationResults  = "chat:conversation-results"   // Atribuição de resultados por conversa
	ChatOperatorQueue        = "chat:operator-queue"         // Fila de espera do operador
	ChatTransferConversation = "chat:transfer-conversation"  // Visualizar e transferir conversas de operadores
	ChatSatisfactionSurvey   = "chat:satisfaction-survey"    // Pesquisa de satisfação dos atendimentos
	ChatSchedule             = "chat:schedule"               // Horário de atendimento e fuso
	ChatUserPermissions      = "chat:user-permissions"       // Gestão de permissão de usuários
	ChatWalletByAttendant    = "chat:wallet-attendant"       // Carteira de clientes por atendente
	ChatWalletVisibility     = "chat:wallet-visibility"      // Configurações de visibilidade de carteira
	ChatSmartTags            = "chat:smart-tags"             // Etiquetas inteligentes
	ChatPauseInvisibleStatus = "chat:pause-invisible-status" // Status em pausa e invisível

	// Automation and Chatbot module features
	AutomationBasicFlows      = "automation:basic-flows"      // Fluxos básicos para recepção e direcionamento de leads em areas ou funis diferentes
	AutomationAPICalls        = "automation:api-calls"        // Chamadas de API e envio de arquivos
	AutomationMessaging       = "automation:messaging"        // Recursos avançados de mensageria (caixa de botões e lista suspensa)
	AutomationPathDivision    = "automation:path-division"    // Divisão de caminhos por propriedades do lead
	AutomationAIAssistants    = "automation:ai-assistants"    // Assistentes de Inteligência artificial
	AutomationSmartQualifying = "automation:smart-qualifying" // Qualificação inteligente

	// Integrations module features
	IntegrationContactsAPI             = "integration:contacts-api"              // API de Criação e busca de contatos
	IntegrationMessagesAPI             = "integration:messages-api"              // API de Envio de mensagens e templates
	IntegrationWalletsAPI              = "integration:wallets-api"               // API de Gerenciamento de carteiras e relatórios
	IntegrationHistoryConversationsAPI = "integration:conversations-history-api" // API de Historico de conversas
	IntegrationCustomFieldsAPI         = "integration:custom-fields-api"         // API de Retorno de campos personalizados
	IntegrationAttendantsAPI           = "integration:attendants-api"            // API de atendimento
	IntegrationMassDeleteContacts      = "integration:mass-delete-contacts"      // API de operação em massa - exclusão de contatos
	IntegrationMassCreateContacts      = "integration:mass-create-contacts"      // API de operação em massa - criação de contatos
	IntegrationMassEmployees           = "integration:mass-employees"            // API de operação em massa de funcionários
	IntegrationDynamicVarsAPI          = "integration:dynamic-vars-api"          // API de personalização de váriáveis dinâmicas em templates
	IntegrationWebhooks                = "integration:webhooks"                  // Criação de webhooks
	IntegrationShopify                 = "integration:shopify"                   // Integração Shopify
	IntegrationNectar                  = "integration:nectar"                    // Integração Néctar
	IntegrationPipedrive               = "integration:pipedrive"                 // Integração Pipedrive
	IntegrationVortexCRM               = "integration:vortex-crm"                // Integração Vórtice CRM
	IntegrationAgendor                 = "integration:agendor"                   // Integração Agendor
	IntegrationExactSales              = "integration:exact-sales"               // Integração Exact Sales
	IntegrationRDCRMCreateDeal         = "integration:rdcrm-create-deal"         // Integração RDCRM: Criar negociação ao iniciar atendimento
	IntegrationRDCRMSendMessages       = "integration:rdcrm-send-messages"       // Integração RDCRM: Enviar mensagens via automação de vendas no CRM
	IntegrationRDCRMStartChat          = "integration:rdcrm-start-chat"          // Integração RDCRM: Botão no CRM para iniciar conversa pelo RDSC
	IntegrationRDCRMOwnerSync          = "integration:rdcrm-owner-sync"          // Integração RDCRM: Sincronização do dono do lead
	IntegrationRDCRMDealUpdate         = "integration:rdcrm-deal-update"         // Integração RDCRM: Atualização da negociação
	IntegrationRDSMCreateLead          = "integration:rdsm-create-lead"          // Integração RDSM: Criar lead no RDSM
	IntegrationRDSMTimelineTracking    = "integration:rdsm-timeline-tracking"    // Integração RDSM: Registrar na timeline do lead abertura e encerramento de chat

	// Analytics module features
	AnalyticsRealTimeMonitoring = "analytics:realtime-monitoring" // Monitoramento em tempo real da equipe
	AnalyticsAttendantDashboard = "analytics:attendant-dashboard" // Dashboard de atendimento (as-is)
	AnalyticsCampaignReports    = "analytics:campaign-reports"    // Relatórios de campanhas
	AnalyticsWordCloud          = "analytics:word-cloud"          // Nuvem de palavras
	AnalyticsAttendantRating    = "analytics:attendant-rating"    // Avaliação de atendimento
	AnalyticsSalesFunnel        = "analytics:sales-funnel"        // Funil de vendas
	AnalyticsCustomDashboards   = "analytics:custom-dashboards"   // Dashboards personalizados
	AnalyticsContactOrigin      = "analytics:contact-origin"      // Origem dos contatos
	AnalyticsEmotionAnalysis    = "analytics:emotion-analysis"    // Análise de emoções
	AnalyticsContactTracking    = "analytics:contact-tracking"    // Trackeamento de origem de contatos

	// ContactManagement module features
	ContactContacts = "contact:contacts" // Gestão de contatos
	ContactTags     = "contact:tags"     // Etiquetas
	ContactBlocking = "contact:blocking" // Bloqueio de contatos
	ContactImport   = "contact:import"   // Importação de contatos

	// Campaigns module features
	CampaignSending              = "campaigns:sending"               // Disparos em massa/campanhas
	CampaignAudienceSegmentation = "campaigns:audience-segmentation" // Segmentação de público
	CampaignScheduling           = "campaigns:scheduling"            // Agendamento de envios
	CampaignSendStatus           = "campaigns:send-status"           // Status de envios
	CampaignWordCloud            = "campaigns:word-cloud"            // Nuvem de palavras

	// Channels module features
	ChannelWhatsapp  = "channels:whatsapp"  // Whatsapp
	ChannelInstagram = "channels:instagram" // Instagram
	ChannelMessenger = "channels:messenger" // Messenger
	ChannelTelegram  = "channels:telegram"  // Telegram
	ChannelEmail     = "channels:email"     // E-mail

	// Additional module features
	// Essas permissões são de limite de valor
	// AdditionalWhatsappNumbers = "additional:whatsapp-numbers" // Números whats app (Chips)
	// AdditionalUsers           = "additional:users"            // Usuários
	// AdditionalChatWidget      = "additional:chat-widget"      // Widget chat
	// AdditionalSupportHistory  = "additional:support-history"  // Historico de atendimento
	// AdditionalFileSending     = "additional:file-sending"     // Envios de arquivos
	// AdditionalAPILimit        = "additional:api-limit"        // Limite de requisições API (entre 300 e 600 req/min)
	// AdditionalAILimit         = "additional:ai-limit"         // Limite de requisições mentor IA
	AdditionalUniqueClients = "additional:unique-clients" // Contador de clientes únicos
)

var (
	AllFeatures = []string{
		ChatChat,
		ChatQuickMessages,
		ChatConversationID,
		ChatRealTimeMonitoring,
		ChatRating,
		ChatConversationResults,
		ChatOperatorQueue,
		ChatTransferConversation,
		ChatSatisfactionSurvey,
		ChatSchedule,
		ChatUserPermissions,
		ChatWalletByAttendant,
		ChatWalletVisibility,
		ChatSmartTags,
		ChatPauseInvisibleStatus,
		AutomationBasicFlows,
		AutomationAPICalls,
		AutomationMessaging,
		AutomationPathDivision,
		AutomationAIAssistants,
		AutomationSmartQualifying,
		IntegrationContactsAPI,
		IntegrationMessagesAPI,
		IntegrationWalletsAPI,
		IntegrationHistoryConversationsAPI,
		IntegrationCustomFieldsAPI,
		IntegrationAttendantsAPI,
		IntegrationMassDeleteContacts,
		IntegrationMassCreateContacts,
		IntegrationMassEmployees,
		IntegrationDynamicVarsAPI,
		IntegrationWebhooks,
		IntegrationShopify,
		IntegrationNectar,
		IntegrationPipedrive,
		IntegrationVortexCRM,
		IntegrationAgendor,
		IntegrationExactSales,
		IntegrationRDCRMCreateDeal,
		IntegrationRDCRMSendMessages,
		IntegrationRDCRMStartChat,
		IntegrationRDCRMOwnerSync,
		IntegrationRDCRMDealUpdate,
		IntegrationRDSMCreateLead,
		IntegrationRDSMTimelineTracking,
		AnalyticsRealTimeMonitoring,
		AnalyticsAttendantDashboard,
		AnalyticsCampaignReports,
		AnalyticsWordCloud,
		AnalyticsAttendantRating,
		AnalyticsSalesFunnel,
		AnalyticsCustomDashboards,
		AnalyticsContactOrigin,
		AnalyticsEmotionAnalysis,
		AnalyticsContactTracking,
		ContactContacts,
		ContactTags,
		ContactBlocking,
		ContactImport,
		CampaignSending,
		CampaignAudienceSegmentation,
		CampaignScheduling,
		CampaignSendStatus,
		CampaignWordCloud,
		ChannelWhatsapp,
		ChannelInstagram,
		ChannelMessenger,
		ChannelTelegram,
		ChannelEmail,
		// AdditionalWhatsappNumbers,
		// AdditionalUsers,
		// AdditionalChatWidget,
		// AdditionalSupportHistory,
		// AdditionalFileSending,
		// AdditionalAPILimit,
		// AdditionalAILimit,
		AdditionalUniqueClients,
	}

	BasicExcluded = []string{
		ChatSmartTags,
		AutomationAPICalls,
		AutomationMessaging,
		AutomationPathDivision,
		AutomationAIAssistants,
		AutomationSmartQualifying,
		IntegrationHistoryConversationsAPI,
		IntegrationCustomFieldsAPI,
		IntegrationAttendantsAPI,
		IntegrationMassDeleteContacts,
		IntegrationMassCreateContacts,
		IntegrationMassEmployees,
		IntegrationDynamicVarsAPI,
		IntegrationWebhooks,
		IntegrationShopify,
		IntegrationNectar,
		IntegrationPipedrive,
		IntegrationVortexCRM,
		IntegrationAgendor,
		IntegrationRDCRMSendMessages,
		IntegrationRDCRMStartChat,
		IntegrationRDCRMOwnerSync,
		IntegrationRDCRMDealUpdate,
		AnalyticsSalesFunnel,
		AnalyticsCustomDashboards,
		AnalyticsContactOrigin,
		AnalyticsEmotionAnalysis,
		AnalyticsContactTracking,
	}

	ProExcluded = []string{
		AutomationAIAssistants,
		AutomationSmartQualifying,
		IntegrationHistoryConversationsAPI,
		IntegrationCustomFieldsAPI,
		IntegrationAttendantsAPI,
		IntegrationMassDeleteContacts,
		IntegrationMassCreateContacts,
		IntegrationMassEmployees,
		IntegrationDynamicVarsAPI,
		AnalyticsCustomDashboards,
		AnalyticsContactTracking,
	}

	BasicFeatures    = GeneratePermissions(AllFeatures, BasicExcluded)
	ProFeatures      = GeneratePermissions(AllFeatures, BasicExcluded)
	AdvancedFeatures = AllFeatures
)

// remove as permissões que o plano não tem de todas
func GeneratePermissions(all, exclude []string) []string {
	excludeMap := make(map[string]bool)
	for _, perm := range exclude {
		excludeMap[perm] = true
	}

	result := []string{}
	for _, perm := range all {
		if !excludeMap[perm] {
			result = append(result, perm)
		}
	}
	return result
}
